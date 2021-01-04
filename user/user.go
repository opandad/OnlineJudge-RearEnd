package user

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	id        int
	email     string
	name      string
	password  string
	authority string
	user_info string //This is json but return string, need special judge.
}

func LoginUseEmail() {
	sqlDB, err := sql.Open("mysql", "online_judge")
	if err != nil {
		fmt.Println("Open connection fail!", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	sqlDB.Ping()
	fmt.Println(gormDB)

	// sqlDB, err := gorm.DB()

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sqlDB.Ping()
}

func Register() {

}

func ForgetPassword() {

}

func LoginUseWechat() {

}
