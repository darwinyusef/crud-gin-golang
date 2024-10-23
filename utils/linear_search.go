package utils

import (
	"github.com/darwinyusef/crud-gin-golang/models"
)

func LinearSearch(users []models.User, targetID string) int {
	for i, user := range users {
		if user.ID == targetID {
			return i
		}
	}
	return -1
}
