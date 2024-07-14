package tests

import (
	"birthday-greeting-service/internal/database"
	"birthday-greeting-service/internal/http/handlers/post"
	"birthday-greeting-service/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeHandler(t *testing.T) {
	database.ConnectDB()                         // подключение к БД
	r := gin.Default()                           // настройка роутера
	r.POST("/employees/", post.SubscribeHandler) //
	w := httptest.NewRecorder()
	employee1 := models.Employee{}
	employee1.ID = 1
	employee1.Name = "Иван" // если сотрудник есть в БД, то код ответа 200
	jsonValue, _ := json.Marshal(employee1)
	req, _ := http.NewRequest(http.MethodPost, "/employees/?subscribe=true", bytes.NewBuffer(jsonValue))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // проверяем, что всё ок
	req, _ = http.NewRequest(http.MethodPost, "/employees/?subscribe=false", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code) // проверяем, что всё ок
	employee2 := models.Employee{}
	employee2.ID = 2
	employee2.Name = "Александр" // если сотрудника нет в БД, то код ответа 404
	jsonValue, _ = json.Marshal(employee2)
	req, _ = http.NewRequest(http.MethodPost, "/employees/?subscribe=false", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code) // проверяем, что всё ок
	req, _ = http.NewRequest(http.MethodPost, "/employees/?subscribe=true", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code) // проверяем, что всё ок
}
