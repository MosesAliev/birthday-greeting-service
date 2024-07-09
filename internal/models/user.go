package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login         string
	Subscriptions []int `gorm:"-"`
}
