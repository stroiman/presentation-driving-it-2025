# Example

```go
func CreateRootHandler() *http.Handler {
}
```

```go
func TestIndexPage(t *testing.T) {
    browser := browser.New(
        browser.WithHandler(CreateRootHandler()),
        browser.WithContext(t.Context()), // Optional
    )
    window, err := browser.Open("https://example.com/")
    assert.NoError(t, err)
    title, err := window.Document().QuerySelector("h1")
    assert.True(t, error)
    assert.NotNil(t, title)
    assert.Equal(t, "Hello, World!", title.TextContent())
}
```

---
layout: top-title
color: slate
---

:: title ::

# Example: A session-based login page

:: content ::

Use tests to drive the implementation of _behaviour_ in the _UI Layer_

- Login session handling
- Redirecting to login when requesting protected resources
- Return to requested resource on successful login

## Strategy

Mock the code validating credentials; Not UI responsibility.

<AdmonitionType type="warning" title="Beware">
This is not a guide to implementing session-based authentication. Many security
concerns are not covered here.
</AdmonitionType>

---

# Coupling to accessibility attributes

Shaman library 

```go
func TestIndexPage(t *testing.T) {

}
```

---

## Embedding V8 in Go

- v8go existed as a project
  - Not all v8 features were supported.

## Mistakes

- Using V8

---
src: ./pages/login-page.md
---

