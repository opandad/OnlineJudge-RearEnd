package database

import (
	"encoding/json"
	"fmt"
	"os"
	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Code     string
	User     string
	Password string
	Database string
}

func ConnectDatabase() {
	filePtr, err := os.Open("../config/database.json")

	//need throw
	if err != nil {
		fmt.Println(err)
		return
	}

	defer filePtr.Close()

	var dbCon DatabaseConfig

	// decoder := json.NewDecoder(filePtr)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&dbCon)
	if err != nil {
		fmt.Println("Decode database config fail!", err.Error())
	} else {
		fmt.Println("Decode database config success!")

		fmt.Println(dbCon)
	}
}
