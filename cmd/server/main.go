package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()      // подключение к БД
	r := router.SetupRouter() // настройка роутера
	r.Run()                   // запуск сервера
}
