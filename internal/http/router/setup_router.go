package router

import (
	"birthday-greeting-service/internal/http/handlers/get"
	"birthday-greeting-service/internal/http/handlers/post"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/employees", post.Auth(get.GetEmployeesHandler))
	r.POST("/employees/", post.Auth(post.SubscribeHandler))
	return r
}
