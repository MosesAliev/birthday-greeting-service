package model

type User struct {
	Login     string      `gorm:"primaryKey"`
	Employees []*Employee `gorm:"many2many:subscriptions;" json:"-"`
}
