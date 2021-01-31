package user

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/web/model"
	"errors"
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
email (string)

@return
成功：返回sessionID, verifyCode(存在session)
失败：返回错误
*/
func SendVerificationCodeToEmailUser(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	verifyCode := verification.RandVerificationCode()
	websocketID := verification.Snowflake()

	//验证邮箱域名是否能ping通，不能ping通则返回错误
	//TODO

	//检查用户是否存在
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return errors.New("mysql数据库连接失败！无法验证用户是否存在。")
	}
	var count int64
	mdb.Model(&model.Email{}).Where("email = ?", websocketInputData.Account).Count(&count)
	fmt.Println(count)
	if count != 0 {
		return errors.New("账号已经被注册，如果忘记密码，请寻回密码")
	}

	//尝试连接websocket数据存放服务器
	rdb, ctx, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return errors.New("redis数据库连接失败！")
	}

	//验证邮箱是否发送正确
	err = email.SendMailByQQ([]string{websocketInputData.Account}, "OnlineJudge", "验证码", verifyCode)
	if err != nil {
		return errors.New("发送邮件验证码失败，请检查邮箱是否正确！")
	}

	//将验证码等数据转换成json
	var userOnlineData model.UserOnlineData
	userOnlineData.VerifyCode = verifyCode
	userOnlineData.WebsocketID = websocketID
	// userOnlineDataJSON, err := json.Marshal(&userOnlineData)
	// if err != nil {
	// 	return errors.New("将用户注册验证码转换成json时失败")
	// }
	// fmt.Println("JSON: ", string(userOnlineDataJSON))

	//验证redis数据库是否加入成功(此处有问题)
	err = rdb.HSet(ctx, websocketInputData.Account, userOnlineData).Err()
	if err != nil {
		return err
		return errors.New("redis数据库加入数据失败")
	}

	websocketOutputData.WebsocketID = websocketID

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
func RegistByEmail(websocketInputData *model.WebsocketInputData) error {
	return nil
}
