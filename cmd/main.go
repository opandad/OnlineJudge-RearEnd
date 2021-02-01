package main

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/server"
)

func main() {
	database.InitMysqlDatabase()
	server.InitServer()
}
