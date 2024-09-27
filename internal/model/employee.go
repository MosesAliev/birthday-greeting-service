package model

import (
	"time"
)

type Employee struct {
	ID    int       `json:"id" gorm:"primaryKey"`
	Name  string    `json:"name"`
	Born  time.Time `json:"born"`
	Users []*User   `gorm:"many2many:subscriptions;" json:"-"`
}
