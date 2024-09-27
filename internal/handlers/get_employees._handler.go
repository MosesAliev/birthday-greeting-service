package handlers

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployeesHandler(c *gin.Context) {
	employees := []model.Employee{}

	database.DB.Db.Find(&employees)

	c.IndentedJSON(http.StatusOK, employees)

}
