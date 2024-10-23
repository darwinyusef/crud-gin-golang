package repositories

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/darwinyusef/crud-gin-golang/models"
)

var users = []models.User{
	{ID: "1", Nombre: "Juan", Apellido: "Pérez", Edad: 30},
	{ID: "2", Nombre: "María", Apellido: "Gómez", Edad: 25},
}

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id string) (models.User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return models.User{}, false
}

func CreateUser(user models.User) {
	users = append(users, user)
}

func UpdateUser(id string, user models.User) bool {
	for i, usr := range users {
		if usr.ID == id {
			users[i] = user
			return true
		}
	}
	return false
}

func SaveUsers(users []models.User) {
	// Implementación para guardar los usuarios
	// Por ejemplo:
	var jsonBytes, err = json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("users.json", jsonBytes, 0644)
}

func DeleteUser(id string) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
