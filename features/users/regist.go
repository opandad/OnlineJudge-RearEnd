package users

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/models"
	"context"
	"fmt"
)

/*
	@Title
	users_manager

	@Description
	与用户登录的模块交互功能

	@Func List
	| func name                      | develop  | unit test |

	|-------------------------------------------------------|

	| RegistByEmail                  |    no    |    no     |

	|SendVerificationCodeToEmailUser |    no    |    no     |
*/

/*
正在开发中

@Title
SendVerificationCodeToEmailUser

@description
发送验证码给用户，并且将发送的验证码储存在session(redis)中，储存的格式为"sessionID" => SessionData

@param
mailAccount (string)

@return
成功或失败 (bool)
*/
func SendVerificationCodeToEmailUser(mailAccount string, ) bool {
	const SENDSUCCESS bool = true

	verifyCode := verification.RandVerificationCode()

	if email.SendMailByQQ([]string{mailAccount}, "OnlineJudge", "验证码", verifyCode) {
		var sessionData models.SessionData = models.SessionData{-1, mailAccount, verifyCode}
		ctx := context.Background()
		sessionID := verification.Snowflake()
		database.ConnectRedisDatabase().SAdd(ctx, string(sessionID), sessionData)

	} else {
		fmt.Println("发送验证码失败，请检查邮箱是否填写正确！")
	}

	return !SENDSUCCESS
}

/*
正在开发中

@Title
RegistByEmail

@description
读取session(redis)中的验证码，验证并返回是否成功注册

@param
email (string)

@return
成功或失败 (bool)
*/
func RegistByEmail() {

}
