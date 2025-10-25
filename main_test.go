package main_test

import (
	. "gost-example"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/shaman"
	"github.com/gost-dom/shaman/ariarole"
	"github.com/stretchr/testify/assert"
)

func TestPrivatePageRedirectsToLogin(t *testing.T) {
	b := browser.New(browser.WithHandler(NewRootHandler(&StubAuthenticator{
		User: User{DisplayName: "Smithy"},
	})))
	win, err := b.Open("http://example.com/private")
	if !assert.NoError(t, err) {
		return
	}

	if !t.Run("Redirects to login page", func(t *testing.T) {
		main := shaman.WindowScope(t, win).Subscope(shaman.ByRole(ariarole.Main))
		title := main.Get(shaman.ByH1)
		assert.Equal(t, "Login", title.TextContent(), "Page title")
		assert.Equal(t, "/login", win.Location().Pathname(), "Location pathname")
		_, hasAlert := main.Query(shaman.ByRole(ariarole.Alert))
		assert.False(t, hasAlert, "Login page has alert on first render")
	}) {
		return
	}
}

func TestLoginWithInvalidCredentials(t *testing.T) {
	b := browser.New(browser.WithHandler(NewRootHandler(&StubAuthenticator{
		Error: ErrInvalidCredentials,
	})))
	win, err := b.Open("http://example.com/login")
	winScope := shaman.WindowScope(t, win)
	assert.NoError(t, err)
	{
		form := winScope.
			Subscope(shaman.ByRole(ariarole.Main)).
			Subscope(shaman.ByRole(ariarole.Form))
		form.Textbox(shaman.ByName("Username")).Write("username")
		form.PasswordText(shaman.ByName("Password")).Write("1234")
		form.Get(shaman.ByRole(ariarole.Button)).Click()
	}

	{
		main := winScope.Subscope(shaman.ByRole(ariarole.Main))
		title := main.Get(shaman.ByH1)
		assert.Equal(t, "Login", title.TextContent())
		assert.Equal(t, "/login", win.Location().Pathname())
		alert := main.Get(shaman.ByRole(ariarole.Alert))
		assert.Contains(t, alert.TextContent(), "Invalid credentials")
	}
}

func TestLoginWithValidCredentials(t *testing.T) {
	b := browser.New(browser.WithHandler(NewRootHandler(&StubAuthenticator{
		User: User{
			DisplayName: "Smithy",
		},
	})))
	win, err := b.Open("http://example.com/login")
	winScope := shaman.WindowScope(t, win)
	assert.NoError(t, err)
	{
		form := winScope.
			Subscope(shaman.ByRole(ariarole.Main)).
			Subscope(shaman.ByRole(ariarole.Form))
		form.Textbox(shaman.ByName("Username")).Write("username")
		form.PasswordText(shaman.ByName("Password")).Write("1234")
		form.Get(shaman.ByRole(ariarole.Button)).Click()
	}

	if !t.Run("Returns to index page", func(t *testing.T) {
		title := shaman.WindowScope(t, win).Get(shaman.ByH1).TextContent()
		assert.Contains(t, title, "Welcome")
		assert.Equal(t, "/", win.Location().Pathname())
	}) {
		return
	}

	t.Run("Allow accessing private resource", func(t *testing.T) {
		win.Navigate("/private")
		title := shaman.WindowScope(t, win).Get(shaman.ByH1)
		assert.Equal(t, "Private Page", title.TextContent())
		assert.Equal(t, "/private", win.Location().Pathname())
	})
}
