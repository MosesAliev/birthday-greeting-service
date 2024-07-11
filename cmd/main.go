package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/handlers/get"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()
	router.GET("/employees", get.GetEmployeesHandler)
	router.Run()
}
