package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/router"
)

func main() {
	database.ConnectDB() // подключение к БД

	r := router.SetupRouter() // настройка роутера

	r.Run() // запуск сервера

}
