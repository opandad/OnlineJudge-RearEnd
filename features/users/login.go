/*
	@Title
	users_manager

	@Description
	与用户登录的模块交互功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| LoginVerifyByEmail  |    no    |    no	 |
	| LoginVerifyByWechat |    no    |    no	 |
	| LoginVerifyPassword |    no    |    no	 |
	| LoginVerifyCode     |    no    |    no	 |
	| FindPasswordByEmail |    no    |    no     |
	| RegistByEmail       |    no    |    no     |
*/
package users

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/models"
	"fmt"
)

/*
正在开发中

@Title
LoginVerifyByEmail

@description
email帐号登录验证模块，其中password传参过来后会进行二次加密进行比较

@param
email, password (string, string)

@return
isSuccess bool
*/
func LoginVerifyByEmail(email string, password string) bool {
	var emails []models.Email
	// var user []models.User
	database.ReconnectDatabase().Debug().Model(&emails).Where("email = ?", email).Association("UserId").Find(&emails[0].User)
	fmt.Println(emails[0])
	// fmt.Println("id:", user[0].ID, "name:", user[0].Name, "password:", user[0].Password, "authority:", user[0].Authority, "userinfo:", user[0].UserInfo)

	//temp var 避免报错
	var isSuccess bool = true

	return isSuccess
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
