package models

type Email struct {
	Email  string `gorm:"primaryKey"`
	UserId int
	User   User `gorm:"foreignKey:UserId"`
}
