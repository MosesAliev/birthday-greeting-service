package database

import (
	"birthday-greeting-service/internal/model"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
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
	projectName := regexp.MustCompile(`^(.*` + "birthday-greeting-service" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	if err := godotenv.Load(string(rootPath) + `/.env`); err != nil {
		log.Print("No .env file found")
		return
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
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

	db.AutoMigrate(&model.Employee{}, &model.User{})

	DB = Dbinstance{
		Db: db,
	}
}
