package users

import (
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/models"

	"github.com/gin-gonic/gin"
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
func SendVerificationCodeToEmailUser(c *gin.Context) {
	var sessionData models.SessionData
	if c.ShouldBind(&sessionData) == nil {
		verifyCode := verification.RandVerificationCode()

		if email.SendMailByQQ([]string{sessionData.Email}, "OnlineJudge", "验证码", verifyCode) {
			// rdb := database.ConnectRedisDatabase()
			// rdb.Set(sessionData.Email, verifyCode)
			c.JSON(200, gin.H{"msg": "验证码发送成功，请到邮箱查收！"})
		} else {
			c.JSON(401, gin.H{"msg": "发送邮件失败，请检查邮箱是否正确！"})
		}
	}
	c.JSON(401, gin.H{"msg": "你在干什么，不要酱紫QAQ！"})
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
func RegistByEmail(c *gin.Context) {

}
