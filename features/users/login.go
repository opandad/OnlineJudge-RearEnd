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
package users

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/models"
	"errors"
)

/*
正在开发中，差加密

@Title
LoginByEmail

@description
email帐号登录验证模块，其中password传参过来后会进行二次加密进行比较

@param
emailAccount, password (string, string)

@return
sessionID, error (int64, error)
*/
func LoginByEmail(websocketInputData *models.WebsocketInputData, websocketOutputData *models.WebsocketOutputData) error {
	var emailAccount models.Email

	//Success query
	mdb, err := database.ReconnectMysqlDatabase()

	if err != nil {
		return errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}

	mdb.Where("email = ?", websocketInputData.Account).Find(&emailAccount)
	mdb.Model(&emailAccount).Association("User").Find(&emailAccount.User)

	//TODO 加密，和数据库比较

	if websocketInputData.Account == emailAccount.Email && websocketInputData.Password == emailAccount.User.Password {
		websocketOutputData.WebsocketID = verification.Snowflake()

		//TODO 记录进redis中，证明已经登录过
		return nil
	} else {
		return errors.New("登录失败，请检查用户名和密码是否正确！")
	}
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
func AutoLogin(sessionID string) {

}
