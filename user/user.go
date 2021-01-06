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
	fmt.Println("login use email")
	db := database.GetDatabaseConnection()

	fmt.Println(db)
}

func Register() {

}

func ForgetPassword() {

}

func LoginUseWechat() {

}
