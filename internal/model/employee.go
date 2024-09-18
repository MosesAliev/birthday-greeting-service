package model

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Born time.Time `json:"-"`
}
