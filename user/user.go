package user

import (
	"OnlineJudge-RearEnd/database"
	"fmt"
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
	db := database.GetDatabaseConnection()

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println(err)
	}
	for true {
		fmt.Println(sqlDB.Ping())
	}

}

func Register() {

}

func ForgetPassword() {

}

func LoginUseWechat() {

}
