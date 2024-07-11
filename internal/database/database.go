package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	// создаём URL для соединения с базой данных.
	// Имя пользователя базы данных, пароль и имя базы данных
	// берутся из переменных окружения,
	// они описаны в файле .env
	dsn := "host=localhost user=postgres password=5280 dbname=birthday-greeting-service port=5432 sslmode=disable TimeZone=Europe/Moscow"
	// создаём подключение к базе данных.
	// В &gorm.Config настраивается логер,
	// который будет сохранять информацию
	// обо всех активностях с базой данных.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")

	DB = Dbinstance{
		Db: db,
	}
}
