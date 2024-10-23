package main

import (
	"github.com/darwinyusef/crud-gin-golang/models"
	"github.com/darwinyusef/crud-gin-golang/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userService := services.UserService{}

	router.GET("/users", func(c *gin.Context) {
		users := userService.GetAllUsers()
		c.JSON(200, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, exists := userService.GetUserByID(id)
		if exists {
			c.JSON(200, user)
		} else {
			c.JSON(404, gin.H{"message": "User not found"})
		}
	})

	router.POST("/users", func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid request body"})
			return
		}
		userService.CreateUser(user)
		c.JSON(201, user)
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid request body"})
			return
		}
		userService := services.UserService{}
		user, exists := userService.GetUserByID(id)
		if exists {
			user.ID = id
			updated := userService.UpdateUser(id, user)
			if updated {
				c.JSON(200, gin.H{"message": "User updated successfully"})
			} else {
				c.JSON(500, gin.H{"message": "Error updating user"})
			}
		} else {
			c.JSON(404, gin.H{"message": "User not found"})
		}
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		userService := services.UserService{}
		_, exists := userService.GetUserByID(id)
		if exists {
			deleted := userService.DeleteUser(id)
			if deleted {
				c.JSON(200, gin.H{"message": "User deleted successfully"})
			} else {
				c.JSON(500, gin.H{"message": "Error deleting user"})
			}
		} else {
			c.JSON(404, gin.H{"message": "User not found"})
		}
	})
}
