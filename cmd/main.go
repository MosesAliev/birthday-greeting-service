package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/handlers/get"
	"birthday-greeting-service/internal/handlers/post"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()
	router.GET("/employees", post.Auth(get.GetEmployeesHandler))
	router.POST("/employees/", post.Auth(post.SubscribeHandler))
	router.Run()
}
