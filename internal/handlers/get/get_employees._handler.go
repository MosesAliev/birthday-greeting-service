package get

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployeesHandler(c *gin.Context) {
	employees := []models.Employee{}
	database.DB.Db.Find(&employees)
	c.IndentedJSON(http.StatusOK, employees)
}
