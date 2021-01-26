<<<<<<< HEAD
package main

import (
	"OnlineJudge-RearEnd/server"
)

func main() {
	// database.InitMysqlDatabase()
	server.InitServer()
}
=======
package main

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/server"
)

func main() {
	database.InitMysqlDatabase()
	server.InitServer()
}
>>>>>>> 11d71640b36e9a9b120394e85a1ecebbd89e1595
