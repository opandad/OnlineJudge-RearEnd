/*
	登录模块之类的功能
*/
package users

import (
	"OnlineJudge-RearEnd/features/database"
	"OnlineJudge-RearEnd/models"
	"fmt"
)

//登录验证模块
func LoginVerifyByEmail() {
	var users []models.User
	database.GetDatabaseConnection().Debug().Find(&users)
	fmt.Println("id:", users[0].ID, "email:", users[0].Email, "name:", users[0].Name, "password:", users[0].Password, "authority:", users[0].Authority, "userinfo:", users[0].UserInfo)
	fmt.Println("id:", users[1].ID, "email:", users[1].Email, "name:", users[1].Name, "password:", users[1].Password, "authority:", users[1].Authority, "userinfo:", users[1].UserInfo)
}

//微信登录
func LoginVerifyByWechat() {

}

//密码验证模块
func LoginVerifyPassword() {

}

//验证码模块
func LoginVerifyCode() {

}

//忘记密码模块
func FindPasswordByEmail() {

}
