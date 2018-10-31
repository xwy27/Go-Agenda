package model

import "errors"

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
	sessionStorage.filePath = "data/session.json"
	return loadSession()
}

func loginStatus() bool {
	return currentUser.Login
}

// 如果未登录，后者不为空
func GetCurrentUserName() (string, error) {
	if err := initSession(); err != nil {
		return "", err
	}
	if loginStatus() {
		return currentUser.CurrentUser, nil
	}
	return "", errors.New("you've not logged in")
}

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

func Logout() error {
	if err := initSession(); err != nil {
		return err
	}
	if loginStatus() {
		currentUser.Login = false
		writeSession()
	}
	return errors.New("you've not logged in")
}

func loadSession() error {
	return sessionStorage.load(&currentUser)
}

func writeSession() error {
	return sessionStorage.write(&currentUser)
}
