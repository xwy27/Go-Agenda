package model

import (
	"errors"
)

// User is a type to store
// the info of any single
// user account
type User struct {
	Username string
	Password string
	Email    string
	Phone    string
}

// usersJSON specific the json type of
// users
type usersJSON struct {
	Users []User
}

// UsersType
type usersType struct {
	storage    Storage
	dictionary map[string]*User
}

var users usersType
var usersDB usersJSON

// InitUsers is the function to initialize Users
func InitUsers() error {
	users.storage.filePath = "../data/users.json"
	users.dictionary = make(map[string]*User)
	load()
	return nil
}

// AddUser 成功返回nil
func AddUser(user *User) error {
	if existedUser := users.dictionary[user.Username]; existedUser != nil {
		return errors.New("username existed")
	}
	users.dictionary[user.Username] = user
	return nil
}

// DeleteUser 成功返回nil
func DeleteUser(user *User) error { return nil }

/*
func FindUsersBy(filter func(*User) bool) ([]User, error) {
	return []User{}, nil
}
*/

// FindUserByName 失败返回nil
func FindUserByName(username string) *User {
	if user, ok := users.dictionary[username]; ok {
		return user
	}
	return nil
}

// CheckPass 成功返回true
func CheckPass(username, password string) bool {
	if user, ok := users.dictionary[username]; ok &&
		user.Password == password {
		return true
	}
	return false
}

func load() {

}

func write() {

}
