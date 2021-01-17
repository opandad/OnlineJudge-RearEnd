package models

type Email struct {
	Email   string `gorm:"primaryKey"`
	UsersId int
}
