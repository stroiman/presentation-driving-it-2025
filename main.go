package main

import (
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

var ErrInvalidCredentials = errors.New("Invalid credentials")

type Authenticator interface {
	Authenticate(username string, password string) (User, error)
}

func main() {
	handler := NewRootHandler(RealAuthenticator{})
	if err := http.ListenAndServe("0.0.0.0:4321", LogMiddleware(handler)); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	} else {
		slog.Info("Server listening on port 4321")
	}
}

type RootHttpHandler struct {
	*http.ServeMux
	Authenticator Authenticator
	loggedIn      bool
}

func NewRootHandler(authenticator Authenticator) *RootHttpHandler {
	handler := RootHttpHandler{
		http.NewServeMux(),
		authenticator,
		false,
	}

	fs := http.FileServer(http.Dir("./static"))
	handler.Handle("/static/", http.StripPrefix("/static/", fs))

	handler.HandleFunc("GET /{$}", handler.GetIndex)
	handler.HandleFunc("GET /private", handler.GetPrivate)
	handler.HandleFunc("GET /login", handler.GetLogin)
	handler.HandleFunc("POST /login", handler.PostLogin)
	return &handler
}

func (h *RootHttpHandler) GetIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate("index.tmpl", w, nil)
}

func (h *RootHttpHandler) GetPrivate(w http.ResponseWriter, r *http.Request) {
	if h.loggedIn {
		renderTemplate("private.tmpl", w, nil)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (h *RootHttpHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	renderTemplate("login.tmpl", w, LoginFormData{})
}

func (h *RootHttpHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	_, err := h.Authenticator.Authenticate(
		getFormValue(r, "username"),
		getFormValue(r, "password"))
	if err == nil {
		h.loggedIn = true
		w.Header().Add("HX-Replace-Url", "/")
		w.Header().Add("HX-Retarget", "body")
		renderTemplate("index.tmpl", w, nil)
	} else {
		renderNamedTemplate("login.tmpl", "login-form-content", w, LoginFormData{
			ErrMsg: "Invalid credentials",
		})
	}
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
	ErrMsg string
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
