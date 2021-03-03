package main

import (
	"OnlineJudge-RearEnd/api/database"
	_ "OnlineJudge-RearEnd/web"
)

func main() {
	database.InitMysqlDatabase()
}
