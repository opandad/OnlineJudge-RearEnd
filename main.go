package main

import (
	"OnlineJudge-RearEnd/database"
	"OnlineJudge-RearEnd/user"
)

func main() {
	database.ConnectDatabase()
	user.LoginUseEmail()
}
