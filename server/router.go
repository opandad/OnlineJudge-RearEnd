package server

import (
	"OnlineJudge-RearEnd/configs"
	"OnlineJudge-RearEnd/models"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	router.Any("/", Websocket)

	// user := router.Group("/user", ping)
	// {
	// 	user.POST("/login_by_email", users.LoginByEmail)
	// 	user.GET("/check_login")
	// 	user.GET("/regist_by_email", users.RegistByEmail)
	// 	user.GET("/get_email_verify_code", users.SendVerificationCodeToEmailUser)
	// 	user.GET("/forget_password_by_email")
	// }

	// problems := router.Group("/problems")
	// {
	// 	problems.GET("/list")
	// 	problems.GET("/problem")
	// 	problems.POST("/submit")
	// }

	// contests := router.Group("/contests")
	// {
	// 	contests.GET("/list")
	// 	contests.GET("/contest")
	// 	contests.GET("/rank")
	// }

	router.Run(configs.REAREND_SERVER_IP + ":" + configs.REAREND_SERVER_PORT)
}

func Router(inputData models.WebsocketInputData) models.WebsocketOutputData {
	var outputData models.WebsocketOutputData
	//登录模块
	// if inputData.Message == "login" {
	// 	if inputData.LoginByWhat == "email" {

	// 	}
	// } else if inputData.Message == "problems" {

	// } else {
	// 	outputData.Message = "404"
	// }

	switch inputData.Message {
	case "login":
		
	default:
		outputData.Message = "404"
	}

	return outputData
}
