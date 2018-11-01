package model

import (
	"errors"
	"os"
)

type Session struct {
	Login       bool
	CurrentUser string
}

var currentUser Session
var sessionStorage Storage
var isSessionInit = false

func initSession() error {
	if isSessionInit {
		return nil
	}
	isSessionInit = true
	sessionStorage.filePath = os.Getenv("GOPATH") + "src/github.com/xwy27/Go-Agenda/data/session.json"
	return loadSession()
}

func loginStatus() bool {
	return currentUser.Login
}

// GetCurrentUserName return the username
// of the user who has logged in. If no
// user has logged in, an error will be
// returned.
func GetCurrentUserName() (string, error) {
	if err := initSession(); err != nil {
		return "", err
	}
	if loginStatus() {
		return currentUser.CurrentUser, nil
	}
	return "", errors.New("you've not logged in")
}

// Login accept a username and a password,
// and it will try to login with these two
// parameters. An error will be returned if
// this attempt fails.
func Login(username, password string) error {
	if err := initSession(); err != nil {
		return err
	}
	if CheckPass(username, password) {
		currentUser.CurrentUser = username
		currentUser.Login = true
		writeSession()
		return nil
	}
	return errors.New("invalid username or password")
}

// Logout tries to logout from current user,
// if you've not logged in, an error will be
// returned.
func Logout() error {
	if err := initSession(); err != nil {
		return err
	}
	if loginStatus() {
		currentUser.Login = false
		writeSession()
		return nil
	}
	return errors.New("you've not logged in")
}

func loadSession() error {
	err := sessionStorage.load(&currentUser)
	if os.IsNotExist(err) {
		currentUser.Login = false
		currentUser.CurrentUser = ""
		return nil
	}
	return err
}

func writeSession() error {
	return sessionStorage.write(&currentUser)
}
