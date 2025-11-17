package models

type User struct {
	CustomModel
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	Role     int    `gorm:"role"`
}

func (*User) TableName() string {
	return "users"
}
