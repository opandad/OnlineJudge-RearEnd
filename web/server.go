package web

import (
	"OnlineJudge-RearEnd/configs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/*
	bug list
	没有掉线机制，容易被ddos
*/

func Init() {
	r := gin.Default()
	r.Any("/", Websocket)

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

	r.Run(configs.REAREND_SERVER_IP + ":" + configs.REAREND_SERVER_PORT)
}

/*
	获取前端接收数据，并返回数据
*/
func Websocket(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		var receiveData FrontEndData

		err = ws.ReadJSON(&receiveData)
		if err != nil {
			fmt.Println("error read json data")
			fmt.Println(err)
			break
		} else {
			err = ws.WriteJSON(Router(receiveData))
			if err != nil {
				fmt.Println("error send json data")
				fmt.Println(err)
				break
			}
		}
	}
}

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
//用户验证统一这里
func Router(receiveData FrontEndData) FrontEndData {
	//检测是否为404，解析请求路径
	var sendData FrontEndData
	var isNot404 bool = false
	var requestPath []string = strings.Split(receiveData.HTTPStatus.RequestPath, "/")

	//test
	fmt.Println("Router output test\n", receiveData, "\n", requestPath)

	//account
	if requestPath[0] == "account" {
		if requestPath[1] == "login" {
			if requestPath[2] == "email" {
				isNot404 = true
				sendData.Data.Email[0].User.ID, sendData.HTTPStatus = receiveData.Data.Email[0].Login(receiveData.WebsocketID)
			}
			if requestPath[2] == "user" {
				isNot404 = true
				_, sendData.HTTPStatus = receiveData.Data.User[0].Login(receiveData.WebsocketID)
			}
		}
		if requestPath[1] == "regist" {
			if requestPath[2] == "email" {
				isNot404 = true
				sendData.Data.Email[0].User.ID, sendData.HTTPStatus = receiveData.Data.Email[0].Regist(receiveData.WebsocketID, receiveData.Data.VerifyCode)
			}
		}
		if requestPath[1] == "userInfo" {

		}
		if requestPath[1] == "verifyCode" {
			if requestPath[2] == "email" {
				isNot404 = true
				sendData.HTTPStatus = receiveData.Data.Email[0].SendVerifyCode()
			}
		}
	}

	//problems
	if requestPath[0] == "problems" {
		if requestPath[1] == "list" {
			isNot404 = true
		}
		if requestPath[1] == "detail" {
			//需要判断题目是否存在，如果不存在返回404
		}
		if requestPath[1] == "submit" {
			//需要验证是否登录
		}
	}

	//404 not found
	if isNot404 == false {
		sendData.HTTPStatus = HTTPStatus{
			Message:     "页面走丢了ToT",
			IsError:     true,
			ErrorCode:   404,
			SubMessage:  "404",
			RequestPath: receiveData.HTTPStatus.RequestPath,
			Method:      receiveData.HTTPStatus.Method,
		}
	}

	return sendData
}
