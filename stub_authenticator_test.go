package main_test

import (
	. "gost-example"
)

type StubAuthenticator struct {
	User  User
	Error error
	Calls [][2]string
}

func (a *StubAuthenticator) Authenticate(username string, password string) (User, error) {
	a.Calls = append(a.Calls, [2]string{username, password})
	return a.User, a.Error
}
