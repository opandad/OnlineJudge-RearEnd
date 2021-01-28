package server

import (
	"OnlineJudge-RearEnd/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Websocket(c *gin.Context) {
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

		var websocketData models.WebsocketInputData

		if json.Unmarshal(message, &websocketData) != nil {
			fmt.Println("前端传输json解析错误：", err)
		}

		message, err = json.Marshal(Router(websocketData))
		if err != nil {
			fmt.Println("后端json解析错误：", err)
		}

		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func notice(c *gin.Context) {

}
