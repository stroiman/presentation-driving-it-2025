---
layout: top-title-two-cols
color: slate
---


:: title ::

<v-switch>
    <template #0><h1>Redirect on private page: Fail</h1></template>
    <template #1><h1>Redirect on private page: Pass</h1></template>
</v-switch>

:: left ::

## Test

```go
func TestPrivatePageRedirectsToLogin(t *testing.T) {
    handler := NewRootHandler(nil)
	b := browser.New(
        browser.WithHandler(handler))
	win, err := b.Open("http://example.com/private")
	assert.NoError(t, err)
	main := shaman.
        WindowScope(t, win).
        Subscope(shaman.ByRole(ariarole.Main))
	title := main.Get(shaman.ByH1)
	assert.Equal(t, "Login", title.TextContent())
	assert.Equal(t, "/login", win.Location().Pathname())
}
```


:: right ::

## Implementation

````md magic-move {at:0}
```go
func NewRootHandler(authenticator Authenticator) *RootHttpHandler {
	handler := RootHttpHandler{
		http.NewServeMux(),
		authenticator,
	}
    // ...
	handler.HandleFunc("GET /private", handler.GetPrivate)
    // ...
}

func (h *RootHttpHandler) GetPrivate(w http.ResponseWriter, r *http.Request) {
	renderTemplate("private.tmpl", w, nil)
}
```
```go
func NewRootHandler(authenticator Authenticator) *RootHttpHandler {
	handler := RootHttpHandler{
		http.NewServeMux(),
		authenticator,
	}
    // ...
	handler.HandleFunc("GET /private", handler.GetPrivate)
    // ...
}

func (h *RootHttpHandler) GetPrivate(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```
````

<v-click hide at="1">

```
--- FAIL: TestPrivatePageRedirectsToLogin (0.02s)
    main_test.go:19:
                Error:          Not equal:
                                expected: "Login"
                                actual  : "Private Page"
                Test:           TestPrivatePageRedirectsToLogin
                Messages:       Page title

```

</v-click>


