package post

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubscribeHandler(c *gin.Context) {
	employee := models.Employee{}
	c.BindJSON(&employee)                 // десериализуем json сотрудника
	res := database.DB.Db.Find(&employee) // находим сотрудника в базе данных
	if res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Нет такого сотрудника"})
		return
	}

	database.DB.Db.Exec("INSERT INTO subscriptions (userlogin, employeeid) VALUES (@userlogin, @employeeid)", sql.Named("userlogin", // добавляем сотрудника в список
		c.GetHeader("login")), sql.Named("employeeid", employee.ID)) // подписок пользователя
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Подписан"})
}
