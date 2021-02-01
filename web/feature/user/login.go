/*
	@Title
	users_manager

	@Description
	与用户登录的模块交互功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| LoginByEmail        |    no    |    no	 |
	| LoginByWechat       |    no    |    no	 |
	| LoginVerifyPassword |    no    |    no	 |
	| LoginVerifyCode     |    no    |    no	 |
	| FindPasswordByEmail |    no    |    no     |
*/
package user

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/web/model"
	"errors"
	"time"
)

/*
正在开发
结果：差加密，差记录

@Title
LoginByEmail

@description
email帐号登录验证模块，其中password传参过来后会进行二次加密进行比较

@param
emailAccount, password (string, string)

@return
sessionID, error (int64, error)
*/
func LoginByEmail(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	var emailAccount model.Email

	//Success query
	mdb, err := database.ReconnectMysqlDatabase()

	if err != nil {
		return errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}

	mdb.Where("email = ?", websocketInputData.Account).Find(&emailAccount)
	mdb.Model(&emailAccount).Association("User").Find(&emailAccount.User)

	//TODO 加密，和数据库比较

	if websocketInputData.Account != emailAccount.Email || websocketInputData.User.Password != emailAccount.User.Password {
		return errors.New("登录失败，请检查用户名和密码是否正确！")
	}

	//TODO 记录进redis中，证明已经登录过
	websocketOutputData.WebsocketID = verification.Snowflake()
	var userOnlineData model.UserOnlineData
	userOnlineData.WebsocketID = websocketOutputData.WebsocketID
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return errors.New("redis数据库连接失败！")
	}

	err = rdb.Set(database.CTX, emailAccount.Email, userOnlineData, time.Minute*30).Err()
	if err != nil {
		return errors.New("redis数据库添加失败！")
	}

	return nil
}

//微信登录
func LoginByWechat() {

}

//密码验证模块
func LoginVerifyPassword() {

}

//验证码模块
func LoginVerifyCode() {

}

//自动登录
func AutoLogin(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) {

}
