package models

type User struct {
	Id       int64 `gorm:"primaryKey"`
	Name     string
	Password string
}
