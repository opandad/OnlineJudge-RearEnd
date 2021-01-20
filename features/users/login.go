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
	var emailAccount models.Email

	//Success query
	database.ReconnectDatabase().Where("email = ?", email).Find(&emailAccount)
	database.ReconnectDatabase().Model(&emailAccount).Association("User").Find(&emailAccount.User)

	var isSuccessLogin bool
	if password == emailAccount.User.Password {
		isSuccessLogin = true
	} else {
		isSuccessLogin = false
	}

	return isSuccessLogin
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
