package main

import (
    "github.com/gin-gonic/gin"
    "github.com/crud-gin-golang/main/services"
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
            c.JSON(400, gin.H{"message": "Invalid