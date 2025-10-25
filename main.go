package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

type Authenticator interface {
	Authenticate(username string, password string) (User, error)
}

func main() {
	handler := NewRootHandler(RealAuthenticator{})
	if err := http.ListenAndServe("0.0.0.0:4321", LogMiddleware(nil, handler)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	} else {
		slog.Info("Server listening on port 4321")
	}
}

type RootHttpHandler struct {
	*http.ServeMux
	Authenticator Authenticator
}

func NewRootHandler(authenticator Authenticator) http.Handler {
	handler := RootHttpHandler{
		http.NewServeMux(),
		authenticator,
	}

	fs := http.FileServer(http.Dir("./static"))
	handler.Handle("/static/", http.StripPrefix("/static/", fs))

	handler.HandleFunc("GET /{$}", handler.GetIndex)
	handler.HandleFunc("GET /private", handler.GetPrivate)
	handler.HandleFunc("GET /login", handler.GetLogin)
	handler.HandleFunc("POST /login", handler.PostLogin)
	return AuthMiddleWare(&handler)
}

func (h *RootHttpHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate("index.tmpl", w, nil)
}

func (h *RootHttpHandler) GetPrivate(w http.ResponseWriter, r *http.Request) {
	_, ok := requestGetAuthenticatedUser(r)
	if ok {
		renderTemplate("private.tmpl", w, nil)
	} else {
		path := fmt.Sprintf("/login?redirectURL=%s", url.QueryEscape(r.URL.RequestURI()))
		http.Redirect(w, r, path, http.StatusSeeOther)
	}
}

func (h *RootHttpHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	redirectURL := r.URL.Query().Get("redirectURL")
	fmt.Println("Render login", redirectURL)
	renderTemplate("login.tmpl", w, LoginFormData{
		RedirectUrl: redirectURL,
	})
}

func (h *RootHttpHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	user, err := h.Authenticator.Authenticate(
		getFormValue(r, "username"),
		getFormValue(r, "password"))
	if err == nil {
		cookie, err = encodeCookie(user)
	}
	if err != nil {
		slog.Error("Error processing login", "err", err)
		renderNamedTemplate("login.tmpl", "login-form-content", w, LoginFormData{
			ErrMsg: "Invalid credentials",
		})
		return
	}
	http.SetCookie(w, cookie)
	redirectUrl := r.PostForm.Get("redirect-url")
	if redirectUrl == "" {
		redirectUrl = "/"
	}
	w.Header().Add("HX-Replace-Url", redirectUrl)
	w.Header().Add("HX-Retarget", "body")
	newReq := requestSetAuthenticatedUser(r, user)
	newReq.Method = "GET"
	newReq.URL.Path = redirectUrl
	newReq.URL.RawQuery = ""
	h.ServeHTTP(w, newReq)
}

func renderTemplate(name string, w http.ResponseWriter, data any) {
	renderNamedTemplate(name, name, w, data)
}

func renderNamedTemplate(fileName, templateName string, w http.ResponseWriter, data any) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("Error parsing template", "err", err)
			renderErrPage(w)
		}
	}()
	w.Header().Set("Content-Type", "text/html")
	t := template.Must(template.ParseGlob("layouts/*.tmpl"))
	t = template.Must(t.ParseFiles(
		fmt.Sprintf("templates/%s", fileName),
	))
	if err := t.ExecuteTemplate(w, templateName, data); err != nil {
		slog.Error("Error executing template", "err", err)
		renderErrPage(w)
	}
}

func renderErrPage(w http.ResponseWriter) {
	w.WriteHeader(500)
	fmt.Fprint(w, `<!doctype html><head><html></head><body><h1>Oh, dear!</h1></body></html>`)
}

type RealAuthenticator struct{}

func (a RealAuthenticator) Authenticate(username string, password string) (User, error) {
	if username != "smithy" || password != "1234" {
		return User{}, ErrInvalidCredentials
	}
	return User{DisplayName: "Smithy"}, nil
}

type User struct {
	DisplayName string
}

type LoginFormData struct {
	RedirectUrl string
	ErrMsg      string
}

func getFormValue(r *http.Request, key string) string {
	// We only need to call this once, but it's idempotent, and the request handler
	// reads a little better when we move the responseibility here.
	r.ParseForm()
	v := r.PostForm[key]
	if len(v) != 1 {
		return ""
	}
	return v[0]
}

var base64Encoding = base64.StdEncoding

func encodeCookie(user User) (*http.Cookie, error) {
	cookie, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("encodeCookig: Error marshalling user: %w", err)
	}
	return &http.Cookie{
		Name:  "auth",
		Value: base64Encoding.EncodeToString(cookie),
	}, nil
}

func decodeCookie(authCookie *http.Cookie) (User, error) {
	var (
		decoded []byte
		user    User
		err     error
	)
	if decoded, err = base64Encoding.DecodeString(authCookie.Value); err != nil {
		err = fmt.Errorf("decodeCookie: base64 decode: %w", err)
	} else if err = json.Unmarshal(decoded, &user); err != nil {
		err = fmt.Errorf("decodeCookie: json unmarshal: %w", err)
	}
	return user, err
}

func AuthMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User
		authCookie, err := r.Cookie("auth")
		if authCookie == nil {
			if err != nil && !errors.Is(err, http.ErrNoCookie) {
				slog.Error("Error reading cookie", "err", err)
			}
		} else {
			if user, err = decodeCookie(authCookie); err == nil {
				r = requestSetAuthenticatedUser(r, user)
			} else {
				slog.Error("Error decoding cookie", "err", err)
			}
		}
		h.ServeHTTP(w, r)
	})
}

type contextKey string

const authenticatedUser contextKey = "authUser"

func requestSetAuthenticatedUser(r *http.Request, user User) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), authenticatedUser, user))
}

func requestGetAuthenticatedUser(r *http.Request) (User, bool) {
	user, ok := r.Context().Value(authenticatedUser).(User)
	return user, ok
}
