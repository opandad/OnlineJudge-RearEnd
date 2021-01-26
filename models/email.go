package models

type Email struct {
	Email  string `gorm:"primaryKey"`
	UserID int
	User   User
}
