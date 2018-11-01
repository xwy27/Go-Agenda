package operation

import (
	"fmt"

	"github.com/xwy27/Go-Agenda/model"
)

// RegisterUser tries to register an account
// with the given information
func RegisterUser(username, password, email, phone string) error {
	return model.AddUser(&model.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone})
}

// LoginUser tries to login with
// the given infomation.
func LoginUser(username, password string) error {
	return model.Login(username, password)
}

// LogoutUser tries to logout
// current user.
func LogoutUser() error {
	return model.Logout()
}

// ListUsers will list all users
// if you've logged in, else an
// error will be returned.
func ListUsers() error {
	_, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	result, err := model.FindUsersBy(func(user *model.User) bool { return true })
	if err != nil {
		return err
	}

	for index, user := range result {
		fmt.Println("==========================================")
		fmt.Printf("User%d:\n", index+1)
		fmt.Println("-Username: " + user.Username)
		fmt.Println("-Email: " + user.Email)
		fmt.Println("-Phone: " + user.Phone)
	}

	return nil
}

// DeleteUser tries to delete the
// account that has logged in, else
// an error will be returned.
func DeleteUser() error {
	username, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}
	err = LogoutUser()
	if err != nil {
		return err
	}
	return model.DeleteUser(username)
}
