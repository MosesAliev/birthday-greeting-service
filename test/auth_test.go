package tests

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

func TestAuth(t *testing.T) {

	database.ConnectDB() // подключение к БД
	r := gin.Default()   // настройка роутера
	r.POST("/auth", handlers.Auth(func(c *gin.Context) {}))
	secretKey := []byte("secret")
	claims := jwt.Map{"login": "a"}
	token, _ := jwt.Sign(jwt.HS256, secretKey, claims)
	req, _ := http.NewRequest(http.MethodPost, "/auth", nil)
	req.Header.Add("Authorization", string(token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}
