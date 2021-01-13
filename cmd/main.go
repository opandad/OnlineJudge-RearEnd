package main

import (
	"OnlineJudge-RearEnd/features/database"
	"OnlineJudge-RearEnd/features/server"
)

func main() {
	database.ConnectDatabase()
	server.InitServer()
}
