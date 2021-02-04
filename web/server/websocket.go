package server

import (
	"OnlineJudge-RearEnd/web/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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
		var websocketData model.WebsocketInputData

		if ws.ReadJSON(&websocketData) != nil {
			fmt.Println("前端传输json解析错误：", err)
		}

		err = ws.WriteJSON(Router(&websocketData))

		if err != nil {
			break
		}
	}
}

func notice(c *gin.Context) {

}
