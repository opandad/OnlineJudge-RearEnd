package users

import (
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/verify_code"
)

/*
	@Title
	users_manager

	@Description
	与用户登录的模块交互功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| RegistByEmail       |    no    |    no     |
*/

/*
正在开发中

@Title
RegistByEmail

@description
注册模块，会根据configs/email.go里面的配置文件发送邮箱配置

@param
email (string)

@return
成功或失败 (bool)
*/
func RegistByEmail(mailAccount string) bool {
	const SENDFAIL bool = false

	if email.SendSingleMailByQQ(mailAccount, "OnlineJudge", "验证码", verify_code.RandVerifyCode("")) {

	}

	// if err != nil {
	// 	fmt.Println("验证码发送失败，错误原因为：", err)
	// 	isRegistSuccess = false
	// }

	return SENDFAIL
}
