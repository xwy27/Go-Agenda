package model

import "errors"

type Session struct {
	Login       bool
	CurrentUser string
}

var CurrentUser Session

func InitSession() error {
	return nil
}

func loginStatus() bool {
	return true
}

// 如果未登录，后者不为空
func GetCurrentUserName() (string, error) {
	return "", errors.New("")
}
