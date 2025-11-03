---
title: Building a Headless Browser in Go
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

drawings:
  persist: true

theme: neversink
---

# Building a Headless Browser in Go

Or: How I Learned To Stop Worrying And Love HTML

by **Peter Strøiman**

::note::

---
layout: top-title
color: slate
---

- [github.com/gost-dom](https://github.com/gost-dom)
- [github.com/stroiman](https://github.com/stroiman)
- [linkedin.com/in/stroiman](https://linkedin.com/in/stroiman)
- [stroiman.photography](https://stroiman.photography)

::title::

# About me

::content::

- Coding since the age of 8
    - Professionally since 1997
- Passionate about Test-Driven Development
- Fine Arts Photographer

---
layout: top-title
color: slate
---

:: title ::

# About the Project, Gost-DOM

:: content ::

A headless browser written in Go

- Written in Go
- To support Test-Driven Development of web applications
- Written in Go
- With a focus on hypermedia frameworks

---
src: ./pages/test-driven-development.md
---

---
src: ./pages/htmx.md
---

---
src: ./pages/example.md
---

---
src: ./pages/building.md
---


## Embedding V8 in Go

- v8go existed as a project
  - Not all v8 features were supported.

## Mistakes

- Using V8

---
src: ./pages/login-page.md
---

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
src: ./pages/features.md
---

---
srd: ./pages/outtro.md
---
