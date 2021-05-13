package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
}

func (*User) TableName() string {
	return "users"
}
