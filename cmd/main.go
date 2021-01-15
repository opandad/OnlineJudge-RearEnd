package main

import (
	"OnlineJudge-RearEnd/database"
	"OnlineJudge-RearEnd/server"
)

func main() {
	database.InitDatabase()
	server.InitServer()
}
