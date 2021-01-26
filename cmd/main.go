package main

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/server"
)

func main() {
	database.InitMysqlDatabase()
	server.InitServer()
}
