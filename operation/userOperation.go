package operation

import (
	"github.com/siskonemilia/Go-Agenda/model"
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

func ListUsers() ([]model.User, error) {
	result, err := model.FindUsersBy(func(user *model.User) bool { return true })
	if err != nil {
		return []model.User{}, err
	}
	return result, err
}

func DeleteUser(username string) error {
	return model.DeleteUser(username)
}
