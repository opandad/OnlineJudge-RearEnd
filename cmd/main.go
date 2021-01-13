package main

import (
	"OnlineJudge-RearEnd/internal/database"
	"OnlineJudge-RearEnd/internal/server"
)

func main() {
	database.ConnectDatabase()
	server.InitServer()
}
