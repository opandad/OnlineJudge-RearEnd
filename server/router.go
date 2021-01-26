package server

import (
	"OnlineJudge-RearEnd/features/users"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	router.GET("/")

	user := router.Group("/user")
	{
		user.GET("/login_by_email", users.LoginByEmail)
		user.GET("/check_login")
		user.GET("/regist_by_email", users.RegistByEmail)
		user.GET("/get_email_verify_code", users.SendVerificationCodeToEmailUser)
		user.GET("forget_password_by_email")
	}
	router.Run()
}
