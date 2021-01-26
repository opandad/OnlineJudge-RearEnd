package server

import (
	"OnlineJudge-RearEnd/features/users"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ping(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func InitServer() {
	router := gin.Default()
	router.GET("/")
	router.GET("/ping", ping)

	user := router.Group("/user")
	{
		user.POST("/login_by_email", users.LoginByEmail)
		user.GET("/check_login")
		user.GET("/regist_by_email", users.RegistByEmail)
		user.GET("/get_email_verify_code", users.SendVerificationCodeToEmailUser)
		user.GET("forget_password_by_email")
	}
	router.Run()
}
