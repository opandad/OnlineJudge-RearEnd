package server

import (
	"OnlineJudge-RearEnd/web/feature/user"
	"OnlineJudge-RearEnd/web/model"
	"fmt"
	"strings"
)

/*
	msg格式：login/xxx/xxx

	route list

	--user
		--login
			--email
			--auto
		--regist
			--email
		--userInfo
		--verifyCode
			--email

	--problems
		--list
		--detail
		--submit
*/
func Router(inputData *model.WebsocketInputData) model.WebsocketOutputData {
	var outputData model.WebsocketOutputData

	//检测是否为404，解析请求路径
	var isRoute bool = false
	var requestPath []string = strings.Split(inputData.RequestPath, "/")

	//test
	fmt.Println("Router output test\n", inputData, "\n", requestPath)

	//login
	if requestPath[0] == "user" {
		if requestPath[1] == "login" {
			if requestPath[2] == "email" {
				isRoute = true
			}
			if requestPath[2] == "auto" {
				isRoute = true
			}
		}
		if requestPath[1] == "regist" {
			if requestPath[2] == "email" {
				isRoute = true
			}

			if requestPath[2] == "verifyCode" {
				if requestPath[3] == "email" {
					isRoute = true
					user.SendVerificationCodeToEmailUser(inputData, &outputData)
				}
			}
		}
		if requestPath[1] == "userInfo" {

		}
	}

	//problems
	if requestPath[0] == "problems" {
		if requestPath[1] == "list" {
			isRoute = true
		}
		if requestPath[1] == "detail" {
			//需要判断题目是否存在，如果不存在返回404
		}
		if requestPath[1] == "submit" {
			//需要验证是否登录
		}
	}

	//404 not found
	if isRoute == false {
		outputData.Message = "404"
	}

	return outputData
}
