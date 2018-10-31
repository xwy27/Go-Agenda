package operation

import (
	"Go-Agenda/model"
	"fmt"
)

func RegisterUser(username, password, email, phone string) error {
	return model.AddUser(&model.User{username, password, email, phone})
}

func LoginUser(username, password string) error {
	return model.Login(username, password)
}

func LogoutUser() error {
	return model.Logout()
}

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
