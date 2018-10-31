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
	dictionary map[string]User
}

var users usersType
var usersDB usersJSON
var isUserInit = false

func initUsers() error {
	if isUserInit {
		return nil
	}
	isUserInit = true
	users.storage.filePath = "data/users.json"
	users.dictionary = make(map[string]User)
	return loadUsers()
}

// AddUser 成功返回nil
func AddUser(user *User) error {
	if err := initUsers(); err != nil {
		return err
	}
	if _, existedUser := users.dictionary[user.Username]; existedUser {
		return errors.New("username existed")
	}
	users.dictionary[user.Username] = User(*user)
	return writeUsers()
}

// DeleteUser 成功返回nil
func DeleteUser(username string) error {
	if err := initUsers(); err != nil {
		return err
	}
	if _, existedUser := users.dictionary[username]; existedUser {
		meetings, err := FindMeetingsBy(func(meeting *Meeting) bool {
			for _, participator := range meeting.Participators {
				if participator.Username == username {
					return true
				}
			}
			return false
		})

		if err != nil {
			return err
		}

		for _, meeting := range meetings {
			err := DeleteParticipator(meeting.Title, username)
			if err != nil {
				return err
			}
		}

		meetings, err = FindMeetingsBy(func(meeting *Meeting) bool {
			if username == meeting.Sponsor {
				return true
			}
			return false
		})

		if err != nil {
			return err
		}

		for _, meeting := range meetings {
			err := DeleteMeeting(meeting.Title)
			if err != nil {
				return err
			}
		}

		delete(users.dictionary, username)
		return writeUsers()
	}
	return errors.New("no such user")
}

func FindUsersBy(filter func(*User) bool) ([]User, error) {
	if err := initUsers(); err != nil {
		return nil, err
	}

	var resultUsers []User
	for _, user := range users.dictionary {
		if filter(&user) {
			resultUsers = append(resultUsers, user)
		}
	}

	return resultUsers, nil
}

// FindUserByName 失败返回nil
func FindUserByName(username string) *User {
	if err := initUsers(); err != nil {
		return nil
	}
	if user, ok := users.dictionary[username]; ok {
		return &user
	}
	return nil
}

// CheckPass 成功返回true
func CheckPass(username, password string) bool {
	if err := initUsers(); err != nil {
		return false
	}
	if user, ok := users.dictionary[username]; ok &&
		user.Password == password {
		return true
	}
	return false
}

func loadUsers() error {
	err := users.storage.load(&usersDB)
	if err != nil {
		return err
	}
	for _, user := range usersDB.Users {
		users.dictionary[user.Username] = User(user)
	}
	return nil
}

func writeUsers() error {
	var newUserDB usersJSON
	for _, user := range users.dictionary {
		newUserDB.Users = append(newUserDB.Users, user)
	}
	return users.storage.write(&newUserDB)
}
