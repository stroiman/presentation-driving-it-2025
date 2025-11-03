# Verify 

A simple interface for credentials validation

```go
type AuthenticatedUserID string

type CredentialsValidator interface {
    Validate(email string, password string) (AuthenticatedUserID, error)
}

type Server struct {
    http.ServeMux
    CredentialsValidator CredentialsValidator
}

func (s *Server) HandlePostLogin(r *http.Request, w http.ResponseWriter) {
}

func NewServer() *Server {
    
}
```

---

```go
type StubbedCredentialsResponse struct {
    User AuthenticatedUser
    Err error
}

func (r StubbedCredentialsResponse) Validate(string,string) (AuthenticatedUser,
error) {
    return User, Err
}
```

---

# Test

```go {*}{maxHeight:'200px'}
func TestLoginPageAuthError(t *testing.T) {
    server := Server{
        Authenticator: StubbedCredentialsResponse{
            Err: errors.New("Stubbed error")
        },
    }
    // Host is ignored, but necessary for cookies to work
    b := browser.New(
        browser.WithServer(server),
    )
    win, err := b.Open("https://example.com/auth/login")
    assert.NoError(t, err)

    doc := win.Document()
    email, err := doc.QuerySelector("input[name='email']")
    assert.NoError(t, err)
    pass, err := doc.QuerySelector("input[name='password']")
    assert.NoError(t, err)
    submit, err := doc.QuerySelector("input[type='submit']")
    assert.NoError(t, err)

    ctrl := KeyboardController{win}
    email.Focus()
    ctrl.SendKeys("user@example.com")
    pass.Focus()
    ctrl.SendKeys("veryS3cre7")
    submit.Click()

    assert.Equal("/auth/login", doc.Location().PathName())
}
```

---
layout: side-title
titlewidth: is-3
align: tl
---

::title::

# Testing Login

::content::

````md magic-move {lines: true, maxHeight:'20px'}
```go {*}{maxHeight:'10px'}
func TestLoginPageAuthError(t *testing.T) {
    server := Server{
        Authenticator: StubbedCredentialsResponse{
            Err: errors.New("Stubbed error")
        },
    }
    // Host is ignored, but necessary for cookies to work
    b := browser.New(
        browser.WithServer(server),
    )
    win, err := b.Open("https://example.com/auth/login")
    assert.NoError(t, err)

    doc := win.Document()
    email, err := doc.QuerySelector("input[name='email']")
    assert.NoError(t, err)
    pass, err := doc.QuerySelector("input[name='password']")
    assert.NoError(t, err)
    submit, err := doc.QuerySelector("input[type='submit']")
    assert.NoError(t, err)

    ctrl := KeyboardController{win}
    email.Focus()
    ctrl.SendKeys("user@example.com")
    pass.Focus()
    ctrl.SendKeys("veryS3cre7")
    submit.Click()

    assert.Equal("/auth/login", doc.Location().PathName())
}
```

```go {*}{maxHeight:'100px'}
func TestLoginPageAuthError(t *testing.T) {
    server := Server{
        Authenticator: StubbedCredentialsResponse{
            Err: errors.New("Stubbed error")
        },
    }
    b := browser.New(
        browser.WithServer(server),
    )
    // Host is ignored, but necessary for cookies to work
    win, err := b.Open("https://example.com/auth/login")
    assert.NoError(t, err)

    form := shaman.WindowScope(t, win)
        .SubScope(shaman.ByRole(ariarole.Main))
        .SubScope(shaman.ByRole(ariarole.Form))

    email := form.Get(
        shaman.ByRole(ariarole.TextBox),
        shaman.ByName("Email"),
    )
    pass := form.Get(
        shaman.ByRole(ariarole.Password),
        shaman.ByName("Password"),
    )
    submit := form.Get(shaman.ByRole(ariarole.Button))

    email.Write("user@example.com")
    pass.Write("veryS3cre7")
    submit.Click()

    assert.Equal("/auth/login", doc.Location().PathName())
}
```

```go
func TestLoginPageAuthError(t *testing.T) {
    server := Server{
        Authenticator: StubbedCredentialsResponse{
            Err: errors.New("Stubbed error")
        },
    }
    b := browser.New(
        browser.WithServer(server),
    )
    // Host is ignored, but necessary for cookies to work
    win, err := b.Open("https://example.com/auth/login") 
    assert.NoError(t, err)

    form := NewLoginForm(win)
    form.Email().Write("user@example.com")
    form.Password().Write("veryS3cre7")
    form.Submit().Click()

    assert.Equal("/auth/login", doc.Location().PathName())
}
```
````
