package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("192.168.121.131:3306", "online_judge_admin:qweasd@/online_judge")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO `online_judge`.`contest_has_problems` (`contest_id`, `problems_id`) VALUES (2, 1);")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT * FROM contest_has_problems")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	_, err = stmtIns.Exec("0", time.Now(), 0, 0, 0, 0)

	// Query the square-number of 13
	err = stmtOut.QueryRow(0).Scan() // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("testtest")
}
