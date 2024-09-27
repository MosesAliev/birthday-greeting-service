package router

import (
	"birthday-greeting-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/employees", handlers.Auth(handlers.GetEmployeesHandler))

	r.POST("/employees/", handlers.Auth(handlers.SubscribeHandler))

	return r
}
