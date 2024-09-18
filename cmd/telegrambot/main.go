package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/handlers"
	"birthday-greeting-service/internal/model"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("employees"),
	),
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Println(err)
		log.Println(os.Getenv("TELEGRAM_APITOKEN"))
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// switch update.Message.Text {
		// case "open":
		msg.ReplyMarkup = numericKeyboard
		// case "close":
		// 	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		// }

		switch update.Message.Text {
		case "employees":
			log.Println("here")
			r := gin.Default() // настройка роутера
			r.GET("/employees", handlers.GetEmployeesHandler)
			req, _ := http.NewRequest(http.MethodGet, "/employees", nil) // запрос на получение списка сотрудников
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			var employees []model.Employee
			json.Unmarshal(w.Body.Bytes(), &employees)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			for _, employee := range employees {
				msg.Text += employee.Name + "\n"
				log.Println(employee.Name)

			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

		}

	}
}
