package users

import (
	"OnlineJudge-RearEnd/models"
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
email (string)

@return
成功：返回sessionID, verifyCode(存在session)
失败：返回错误
*/
func SendVerificationCodeToEmailUser(websocketInputData *models.WebsocketInputData, websocketOutputData *models.WebsocketOutputData) error {
	// verifyCode := verification.RandVerificationCode()

	// err := email.SendMailByQQ([]string{websocketInputData.Account}, "OnlineJudge", "验证码", verifyCode)

	// //验证邮箱是否发送正确
	// if err == nil {
	// 	websocketOutputData.SessionID = verification.Snowflake()
	// 	rdb, ctx, err := database.ConnectRedisDatabase()

	// 	//验证redis是否连接成功
	// 	if err == nil {
	// 		rdb.Set(ctx, websocketOutputData.SessionID, verifyCode, 1000)

	// 		//TODO
	// 		return nil
	// 	}

	// 	return errors.New("redis数据库连接失败！")
	// } else {
	// 	return errors.New("发送邮件验证码失败，请检查邮箱是否正确！")
	// }
	return nil
}

/*
正在开发中

@Title
RegistByEmail

@description
读取session(redis)中的验证码，验证并返回是否成功注册

@param
email, password, verifyCode (string, string, string)

@return
成功或失败 (bool)
*/
func RegistByEmail(websocketInputData *models.WebsocketInputData) error {
	return nil
}
