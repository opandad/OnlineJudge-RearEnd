package server

import (
	"OnlineJudge-RearEnd/web/model"
	"strings"
)

/*
	msg格式：login/xxx/xxx
*/
func Router(inputData *model.WebsocketInputData) model.WebsocketOutputData {
	var outputData model.WebsocketOutputData

	//检测是否为404，解析请求路径
	var isRoute bool = false
	var requestPath []string = strings.Split(inputData.RequestPath, "/")

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
