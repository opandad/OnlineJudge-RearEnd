package main

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web"
)

func main() {
	database.InitMysqlDatabase()
	web.Init()
}
