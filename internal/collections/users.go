package collections

import (
	"errors"
	"fmt"
	"xuoxod/adminhelper/internal/models"
	"xuoxod/adminhelper/internal/utils"
)

type iUsers interface {
	AddUser(models.User)
	RemoveUserByEmail(email string) (bool, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers() map[string]models.User
}

type Users struct {
	Users map[string]models.User `json:"users"`
}

func InitUsers() map[string]models.User {
	Users := make(map[string]models.User)
	return Users
}

func (u *Users) AddUser(user models.User) {
	id, err := utils.GenerateUserDefinedRandomNumber(514, 8007)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		user.SetID(fmt.Sprintf("%d", id))
	}

	u.Users[user.GetEmail()] = user
}

func (u *Users) RemoveUserByEmail(email string) (bool, error) {
	_, ok := u.Users[email]

	if !ok {
		return false, fmt.Errorf(fmt.Sprintf("User %s does not exist\n", email))
	}
	delete(u.Users, email)
	return true, nil
}

func (u *Users) GetUserByEmail(email string) (models.User, error) {
	user, ok := u.Users[email]

	if !ok {
		fmt.Printf("User %s not found\n", email)
		return models.User{}, errors.New("User name found")
	}

	return user, nil
}

func (u *Users) GetUsers() map[string]models.User {
	return u.Users
}

func NewUsers() iUsers {
	return &Users{
		Users: make(map[string]models.User),
	}
}
