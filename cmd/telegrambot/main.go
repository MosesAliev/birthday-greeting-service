package main

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/handlers"
	"birthday-greeting-service/internal/model"
	"birthday-greeting-service/internal/router"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var keyboard = tgbotapi.NewReplyKeyboard(
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

		panic(err)

	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	var chatIDs []int64
	go func() {
		for {
			now := time.Now()

			year, month, day := now.Date()

			r := gin.Default() // настройка роутера

			r.GET("/employees", handlers.GetEmployeesHandler)

			req, _ := http.NewRequest(http.MethodGet, "/employees", nil) // запрос на получение списка сотрудников

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			var employees []model.Employee
			json.Unmarshal(w.Body.Bytes(), &employees)

			time.Sleep(24 * time.Hour)

			for _, employee := range employees {
				if year == now.Year() && month == now.Month() && day == now.Day() {
					if len(chatIDs) != 0 {
						for _, chatID := range chatIDs {
							msg := tgbotapi.NewMessage(chatID, "")

							msg.Text = fmt.Sprintf("У %s день рождения", employee.Name)

							if _, err := bot.Send(msg); err != nil {
								log.Panic(err)

							}

						}

					}

				}

			}

		}

	}()

	for update := range updates {
		if update.Message == nil { // игнорировать обновления, не относящиеся к сообщениям
			if update.CallbackQuery != nil {
				if update.CallbackQuery.Data == "Подписаться" {
					chatIDs = append(chatIDs, update.CallbackQuery.From.ID)

					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Подписаться")

					r := router.SetupRouter() // настройка роутера

					data := strings.Split(update.CallbackQuery.Message.Text, " ")

					employeeID, _ := strconv.Atoi(data[0])

					chatID := strconv.Itoa(int(update.CallbackQuery.From.ID))

					res := database.DB.Db.Create(&model.User{Login: chatID})

					if res.Error != nil {
						log.Println("пользователь уже в системе")

					}

					token := jwt.NewWithClaims(jwt.SigningMethodHS256,
						jwt.MapClaims{
							"login": chatID,
						})

					secretKey := []byte("auth")

					signedToken, err := token.SignedString(secretKey)

					if err != nil {
						log.Println(err)

					}

					signedToken = "token " + signedToken
					subscribeRequest, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/employees/?subscribe=true&id=%d", employeeID), nil)

					subscribeRequest.Header.Add("Authorization", signedToken)

					w := httptest.NewRecorder()

					r.ServeHTTP(w, subscribeRequest)

					response := model.Response{}

					json.Unmarshal(w.Body.Bytes(), &response)

					msg.Text = response.Message
					if _, err := bot.Send(msg); err != nil {
						panic(err)

					}

					continue
				} else if update.CallbackQuery.Data == "Отписаться" {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Отписаться")

					r := router.SetupRouter() // настройка роутер

					data := strings.Split(update.CallbackQuery.Message.Text, " ")

					employeeID, _ := strconv.Atoi(data[0])

					chatID := strconv.Itoa(int(update.CallbackQuery.From.ID))

					res := database.DB.Db.Create(&model.User{Login: chatID})

					if res.Error != nil {
						log.Println("пользователь уже в системе")

					}

					token := jwt.NewWithClaims(jwt.SigningMethodHS256,
						jwt.MapClaims{
							"login": chatID,
						})

					secretKey := []byte("auth")

					signedToken, err := token.SignedString(secretKey)

					if err != nil {
						log.Println(err)

					}

					signedToken = "token " + signedToken
					subscribeRequest, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/employees/?subscribe=false&id=%d", employeeID), nil)

					subscribeRequest.Header.Add("Authorization", signedToken)

					w := httptest.NewRecorder()

					r.ServeHTTP(w, subscribeRequest)

					response := model.Response{}

					json.Unmarshal(w.Body.Bytes(), &response)

					msg.Text = response.Message
					if _, err := bot.Send(msg); err != nil {
						panic(err)

					}

					continue
				}

				// Respond to the callback query, telling Telegram to show the user
				// a message with the data received.
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")

				if _, err := bot.Request(callback); err != nil {
					panic(err)

				}

				// And finally, send a message containing the data received.
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

				var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Подписаться", "Подписаться"),
					tgbotapi.NewInlineKeyboardButtonData("Отписаться", "Отписаться"),
				))

				msg.ReplyMarkup = inlineKeyboard
				if _, err := bot.Send(msg); err != nil {
					panic(err)

				}

				continue
			}

			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		msg.ReplyMarkup = keyboard
		switch update.Message.Text {
		case "employees":
			chatID := strconv.Itoa(int(update.Message.Chat.ID))

			res := database.DB.Db.Create(&model.User{Login: chatID})

			if res.Error != nil {
				log.Println("пользователь уже в системе")

			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256,
				jwt.MapClaims{
					"login": chatID,
				})

			secretKey := []byte("auth")

			signedToken, err := token.SignedString(secretKey)

			if err != nil {
				log.Println(err)

			}

			signedToken = "token " + signedToken
			r := router.SetupRouter() // настройка роутера

			req, _ := http.NewRequest(http.MethodGet, "/employees", nil) // запрос на получение списка сотрудников

			req.Header.Add("Authorization", signedToken)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			var employees []model.Employee
			json.Unmarshal(w.Body.Bytes(), &employees)

			var buttons []tgbotapi.InlineKeyboardButton
			for _, employee := range employees {
				buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d %s", employee.ID, employee.Name),
					fmt.Sprintf("%d %s %s", employee.ID, employee.Name, employee.Born.Format("2006-01-02"))))

			}

			var newInlineKeyboardRows [][]tgbotapi.InlineKeyboardButton
			for i, button := range buttons {
				newInlineKeyboardRows = append(newInlineKeyboardRows, tgbotapi.NewInlineKeyboardRow())

				newInlineKeyboardRows[i] = append(newInlineKeyboardRows[i], button)

			}

			var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(newInlineKeyboardRows...)

			msg.Text = "Сотрудники"
			msg.ReplyMarkup = inlineKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)

			}

		case "/start":
			msg.Text = "/start"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)

			}

		default:
			messageTextSplit := strings.Split(update.Message.Text, " ")

			employeeID, err := strconv.Atoi(messageTextSplit[0])

			if err != nil {
				msg.Text = "Неправильный запрос"
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)

				}

			} else {
				chatIDs = append(chatIDs, update.CallbackQuery.From.ID)

				r := router.SetupRouter() // настройка роутера

				chatID := strconv.Itoa(int(update.CallbackQuery.From.ID))

				res := database.DB.Db.Create(&model.User{Login: chatID})

				if res.Error != nil {
					log.Println("пользователь уже в системе")

				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256,
					jwt.MapClaims{
						"login": chatID,
					})

				secretKey := []byte("auth")

				signedToken, err := token.SignedString(secretKey)

				if err != nil {
					log.Println(err)

				}

				signedToken = "token " + signedToken
				subscribeRequest, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/employees/?subscribe=true&id=%d", employeeID), nil)

				subscribeRequest.Header.Add("Authorization", signedToken)

				w := httptest.NewRecorder()

				r.ServeHTTP(w, subscribeRequest)

				response := model.Response{}

				json.Unmarshal(w.Body.Bytes(), &response)

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)

				}

			}

		}

	}

}
