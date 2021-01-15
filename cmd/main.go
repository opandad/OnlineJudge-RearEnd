package main

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/server"
)

func main() {
	database.InitDatabase()

	server.InitServer()
}
