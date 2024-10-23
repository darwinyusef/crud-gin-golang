package services

import (
	"github.com/crud-gin-golang/main/models"
	"github.com/crud-gin-golang/main/repositories"
	"github.com/crud-gin-golang/main/utils"
)

type UserService struct{}

func (us UserService) GetAllUsers() []models.User {
	return repositories.GetAllUsers()
}

func (us UserService) GetUserByID(id string) (models.User, bool) {
	return repositories.GetUserByID(id)
}

func (us UserService) CreateUser(user models.User) {
	repositories.CreateUser(user)
}

func (us UserService) UpdateUser(id string, user models.User) bool {
	return repositories.UpdateUser(id, user)
}

func (us UserService) DeleteUser(id string) bool {
	return repositories.DeleteUser(id)
}

func (us UserService) FindUserByID(id string) (models.User, bool) {
	users := repositories.GetAllUsers()
	index := utils.LinearSearch(users, id)
	if index != -1 {
		return users[index], true
	}
	return models.User{}, false
}
