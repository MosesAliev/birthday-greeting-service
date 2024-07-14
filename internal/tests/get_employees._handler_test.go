package tests

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/http/handlers/get"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployeesHandler(t *testing.T) {
	database.ConnectDB()                                         // подключение к БД
	r := gin.Default()                                           // настройка роутера
	r.GET("/employees", get.GetEmployeesHandler)                 //
	req, _ := http.NewRequest(http.MethodGet, "/employees", nil) // запрос на получение списка сотрудников
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // проверяем, что всё ок
}
