package services

import (
	"github.com/darwinyusef/crud-gin-golang/models"
	"github.com/darwinyusef/crud-gin-golang/repositories"
	"github.com/darwinyusef/crud-gin-golang/utils"
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
	users := repositories.GetAllUsers()
	index := utils.LinearSearch(users, id)
	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		repositories.SaveUsers(users)
		return true
	}
	return false
}

func (us UserService) FindUserByID(id string) (models.User, bool) {
	users := repositories.GetAllUsers()
	index := utils.LinearSearch(users, id)
	if index != -1 {
		return users[index], true
	}
	return models.User{}, false
}
