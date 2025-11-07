---
layout: section
color: slate
---

# Example - Session-based Authentication

---
layout: top-title
color: slate
---

<AdmonitionType type="warning" title="Beware" v-drag="[543,327,337,133,-13]">
This is not a guide to implementing session-based authentication. Many security
concerns are not covered here.

This was chosen as an example many web developers should be familiar with.
</AdmonitionType>

:: title ::

# Example - Desired Behaviour

:: content ::

Implementing username/password login flow

## Behaviour in the UI layer

- Requesting a private page redirects to login if unauthenticated
- Submitting invalid displays an error message
- Submitting valid credentials redirects to the previous page.
  - Redirect to index page, if login opened directly.

---
layout: top-title
color: slate
---

<v-click>
<ArrowDraw v-drag="[470,173,174,42,170]"/>

<StickyNote title="Go 101" v-drag="[599,131,227,125]">

An "embedded field". The `RootHttpHandler` exposes all public methods on the
ServeMux, making it a valid http `Handler` implementation.

</StickyNote>
</v-click>

<v-click>
<ArrowDraw v-drag="[399,314,140,52,213]"/>

<StickyNote color="red" v-drag="[530,356,249,103]">

A simplification for the sake of the example; a real application would have some
tooling to help build the dependency graph at startup.

</StickyNote>

</v-click>

:: title ::

# Root HTTP handler

:: content ::

```go
type Authenticator interface {
	Authenticate(username string, password string) (User, error)
}


type RootHttpHandler struct {
	*http.ServeMux // Usually called a "router" in other libraries
	Authenticator Authenticator
}

func NewRootHandler(authenticator Authenticator) *RootHttpHandler {
	handler := RootHttpHandler{
		http.NewServeMux(),
		authenticator,
	}
    // Configure the serveMux
    return &handler
}
```

---
src: ./example-login-page.md
---


<!-- --- -->
<!-- src: ./example-redirect.md -->
<!-- --- -->

---
layout: top-title
color: slate
---

:: title ::

# Handle failed login - `Authenticator` Stub

:: content ::

```go
type StubAuthenticator struct {
	User  User
	Error error
	Calls [][2]string
}

func (a *StubAuthenticator) Authenticate(username string, password string) (User, error) {
	a.Calls = append(a.Calls, [2]string{username, password})
	return a.User, a.Error
}
```

---
layout: top-title-two-cols
color: slate
columns: is-6-6
---


:: title ::

<v-switch>
    <template #0><h1>Login With Invalid Credentials: Fail</h1></template>
    <template #1><h1>Login With Invalid Credentials: Pass</h1></template>
    <template #2><h1>Login With Invalid Credentials: Pass</h1></template>
</v-switch>

:: left ::

## Test

```go {all|12-23|27-39}{at: 0,maxHeight:'23rem'}
func TestLoginWithInvalidCredentials(t *testing.T) {
	auth := &StubAuthenticator{
		Error: ErrInvalidCredentials,
	}
	handler := NewRootHandler(auth)
	b := browser.New(browser.WithHandler(handler))
	win, err := b.Open("http://example.com/login")
	winScope := shaman.WindowScope(t, win)
	assert.NoError(t, err)

	{
		form := winScope.
			Subscope(shaman.ByRole(ariarole.Main)).
			Subscope(shaman.ByRole(ariarole.Form))
		form.
			Textbox(shaman.ByName("Username")).
			Write("username")
		form.
			PasswordText(shaman.ByName("Password"))
			.Write("1234")
		form.
			Get(shaman.ByRole(ariarole.Button)).
			Click()
	}

	{
		main := shaman.
            WindowScope(t, win).
            Subscope(shaman.ByRole(ariarole.Main))
		title := main.Get(shaman.ByH1)
		assert.Equal(t, "Login", title.TextContent())
		assert.Equal(t, 
            "/login", 
            win.Location().Pathname())
		alert := main.
            Get(shaman.ByRole(ariarole.Alert))
		assert.Contains(t, 
            alert.TextContent(), 
            "Invalid credentials")
	}
}
```

:: right ::

## Handler

````md magic-move {at:0,maxHeight:'10px'}
```go
func NewRootHandler(
    authenticator Authenticator,
) *RootHttpHandler {
	// ...
	handler.HandleFunc("GET /login", handler.GetLogin)
	return &handler
}
```

```go
func NewRootHandler(
    authenticator Authenticator,
) *RootHttpHandler {
	// ...
	handler.HandleFunc("GET /login", handler.GetLogin)
	handler.HandleFunc("POST /login", handler.PostLogin)
	return &handler
}

func (h *RootHttpHandler) PostLogin(
	w http.ResponseWriter, 
	r *http.Request) {
	renderTemplate("login.tmpl", w, LoginFormData{
		ErrMsg: "Invalid credentials",
	})
}
```
````

## Template

````md magic-move {at:0}
```html
  <div class="form-field-list">
    ...
    <button type="submit" class="btn-cta">
        Submit
    </button>
  </div>
```

```html
  <div class="form-field-list">
    ...
    <button type="submit" class="btn-cta">
        Submit
    </button>
    {{ if (ne .ErrMsg "") }}
      <div class="alert" role="alert">
        {{ .ErrMsg }}
      </div>
    {{ end }}
  </div>
```
````
