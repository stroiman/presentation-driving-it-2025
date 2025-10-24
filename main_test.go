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
	b := browser.New(browser.WithHandler(NewRootHandler(nil)))
	win, err := b.Open("http://example.com/private")
	assert.NoError(t, err)
	main := shaman.WindowScope(t, win).Subscope(shaman.ByRole(ariarole.Main))
	title := main.Get(shaman.ByH1)
	assert.Equal(t, "Login", title.TextContent(), "Page title")
	assert.Equal(t, "/login", win.Location().Pathname(), "Location pathname")
}
