---
layout: side-title
titlewidth: is-3
align: tl
dragPos:
  "708": 781,144,171,114,-6
---

::title::

# Testing Login

::content::

````md magic-move {lines: true}
```go {*}{maxHeight:'10px'}
func TestLoginPageAuthError(t *testing.T) {
    auth := StubbedCredentialsResponse{
        Err: errors.New("Stubbed error"),
    }
    server := NewServer(auth)
    b := browser.New(browser.WithHandler(server))
    // Host is ignored, but necessary for cookies to work
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
    auth := StubbedCredentialsResponse{
        Err: errors.New("Stubbed error"),
    }
    server := NewServer(auth)
    b := browser.New(browser.WithHandler(server))
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
    auth := StubbedCredentialsResponse{
        Err: errors.New("Stubbed error"),
    }
    server := NewServer(auth)
    b := browser.New(browser.WithHandler(server))
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

<ArrowDraw v-drag="[611,77,140,52,158]" v-click.hide="1"/>
<ArrowDraw v-drag="[224,191,140,52,-18]" v-click.hide="1"/>

<StickyNote v-drag="[31,244,257,97]" v-click.hide="1">

Host name is ignored, and the page loads with just `Open("/auth/login")`, but
cookies don't work without a host.

</StickyNote>

<StickyNote v-drag="[749,17,216,114]" v-click.hide="1">

`WithHandler` bypass the TCP stack, connecting outgoing HTTP requests directly
to the root `http.Handler`.


</StickyNote>


<div v-click="[1,2]">

<ArrowDraw v-drag="[588,340,140,52,205]" />
<StickyNote v-drag="[732,372,180,105,-7]">

"Name" refers to "Accessibility Name", which is the "label" for an input.

</StickyNote>

<StickyNote v-drag="[40,154,180,180,6]">

The `shaman` library helps write tests coupled to accessibility roles reflecting
how a user sees the application.

</StickyNote>

<StickyNote v-drag="708,195,196,198">

Searching inside `main` ARIA role helps ensure better accessibility through
landmark navigation.

</StickyNote>

<ArrowDraw v-drag="[682,200,108,52,166]"/>
</div>

<div v-click="2">
<StickyNote v-drag="[730,327,180,180]">

A `LoginForm` test helper reduces code duplication for more concise tests more
clearly communicating _intent_.

</StickyNote>
</div>
