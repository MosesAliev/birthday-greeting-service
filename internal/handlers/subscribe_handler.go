package handlers

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SubscribeHandler(c *gin.Context) {
	param, _ := c.GetQuery("subscribe")
	subscribe, err := strconv.ParseBool(param)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "неверный запрос"})
		return
	}

	employee := model.Employee{}

	c.BindJSON(&employee)                  // десериализуем json сотрудника
	res := database.DB.Db.First(&employee) // находим сотрудника в базе данных

	if res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Нет такого сотрудника"})
		return
	}
	if subscribe {
		res := database.DB.Db.Exec("INSERT INTO subscriptions (userlogin, employeeid) VALUES (@userlogin, @employeeid)", sql.Named("userlogin", // добавляем сотрудника в список
			c.GetHeader("login")), sql.Named("employeeid", employee.ID)) // подписок пользователя
		if res.Error != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Вы уже подписаны"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Подписка"})
	} else {
		res := database.DB.Db.Exec("DELETE FROM subscriptions WHERE userlogin=@userlogin AND employeeid=@employeeid", sql.Named("userlogin", // добавляем сотрудника в список
			c.GetHeader("login")), sql.Named("employeeid", employee.ID)) // подписок пользователя
		if res.Error != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Сотрудник не найден в ваших подписках"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Отписка"})
	}

}
