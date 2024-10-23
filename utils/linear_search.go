package utils

import (
	"github.com/crud-gin-golang/main/models"
)

func LinearSearch(users []models.User, targetID string) int {
	for i, user := range users {
		if user.ID == targetID {
			return i
		}
	}
	return -1
}
