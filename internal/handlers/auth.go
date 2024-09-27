package handlers

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(next func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization"), " ")[1]
		secretKey := []byte("auth")

		jwtToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		log.Println(c.GetHeader("Autorization"))

		res, ok := jwtToken.Claims.(jwt.MapClaims)

		// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
		// если Сlaims вдруг оказжется другого типа, мы получим панику
		if !ok {
			log.Printf("failed to typecast to jwt.MapCalims")

			return
		}

		loginRaw := res["login"]
		login, ok := loginRaw.(string)

		if !ok {
			log.Printf("failed to typecast to string login")

			return
		}

		// обратите внимание, что при создании мы указывали тип []string, однако тут приводим к []inteface{}
		// так происходит, потому что json не строго типизированный, из-за чего при парсинге нельзя точно
		// определить тип слайса.
		user := model.User{}

		database.DB.Db.First(&user, "login = ?", login)

		if len(user.Login) == 0 {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Не авторизован"})

			return
		}

		c.Request.Header.Add("login", user.Login)

		next(c)

	}

}
