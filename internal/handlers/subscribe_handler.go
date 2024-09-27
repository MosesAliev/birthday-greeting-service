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

	id, _ := c.GetQuery("id")

	subscribe, err := strconv.ParseBool(param)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "неверный запрос"})

		return
	}

	employee := model.Employee{}

	// c.BindJSON(&employee) // десериализуем json сотрудника
	// res := database.DB.Db.First(&employee) // находим сотрудника в базе данных
	employee_id, _ := strconv.Atoi(id)

	res := database.DB.Db.Where("id = ?", employee_id).First(&employee)

	if res.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Нет такого сотрудника"})

		return
	}

	if subscribe {
		res := database.DB.Db.Exec("INSERT INTO subscriptions (user_login, employee_id) VALUES (@user_login, @employee_id)", sql.Named("user_login", // добавляем сотрудника в список
			c.GetHeader("login")), sql.Named("employee_id", employee.ID)) // подписок пользователя

		if res.Error != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Вы уже подписаны"})

			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Подписка"})

	} else {
		res := database.DB.Db.Exec("DELETE FROM subscriptions WHERE user_login=@user_login AND employee_id=@employee_id", sql.Named("user_login", // добавляем сотрудника в список
			c.GetHeader("login")), sql.Named("employee_id", employee.ID)) // подписок пользователя

		if res.Error != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Сотрудник не найден в ваших подписках"})

			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Отписка"})

	}

}
