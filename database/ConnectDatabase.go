package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	dsn := "online_judge_admin:qweasd@tcp(10.18.121.241:3306)/online_judge?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(db)
	}
}
