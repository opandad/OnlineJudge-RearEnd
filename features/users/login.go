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
	var users []models.User
	database.ReconnectDatabase().Debug().Where("email = ?", email).First(&users)
	fmt.Println("id:", users[0].ID, "name:", users[0].Name, "password:", users[0].Password, "authority:", users[0].Authority, "userinfo:", users[0].UserInfo)

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

//忘记密码模块
func FindPasswordByEmail() {

}

/*
正在开发中

@Title
RegistByEmail

@description
注册模块，会根据configs/email.go里面的配置文件发送邮箱配置

@param
null

@return
null
*/
func RegistByEmail() {

}
