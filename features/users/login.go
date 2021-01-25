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
	"OnlineJudge-RearEnd/models"

	"github.com/gin-gonic/gin"
)

/*
正在开发中

@Title
LoginByEmail

@description
email帐号登录验证模块，其中password传参过来后会进行二次加密进行比较

@param
email, password (string, string)

@return
isSuccess bool
*/
func LoginByEmail(c *gin.Context) {
	var sessionData models.SessionData
	if c.ShouldBind(&sessionData) == nil {
		var emailAccount models.Email

		//Success query
		database.ReconnectMysqlDatabase().Where("email = ?", sessionData.Email).Find(&emailAccount)
		database.ReconnectMysqlDatabase().Model(&emailAccount).Association("User").Find(&emailAccount.User)

		if sessionData.Email == emailAccount.Email && sessionData.Password == emailAccount.User.Password {
			c.JSON(200, gin.H{
				"userID":    emailAccount.User.ID,
				"authority": emailAccount.User.Authority,
				"msg":       "登录成功！",
			})
		} else {
			c.JSON(401, gin.H{"msg": "用户名或密码错误，请重新登录！"})
		}
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
