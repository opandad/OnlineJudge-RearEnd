package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	ServerIP       string `json:"ServerIP"`
	Port           string `json:"ServerPort"`
	User           string `json:"User"`
	Password       string `json:"Password"`
	Database       string `json:"Database"`
	DatabaseEncode string `json:"DatabaseEncode"`
}

func ConnectDatabase() {
	file, err := os.Open("./config/database.json")
	if err != nil {
		fmt.Println("Open database config file fail!", err.Error())
	} else {
		fmt.Println("Open database config file success!")
	}
	defer file.Close()

	var dbConfig DatabaseConfig
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal([]byte(byteValue), &dbConfig)

	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.ServerIP + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=" + dbConfig.DatabaseEncode + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect database fail. Please check database connect config!")
		fmt.Println(err)
	} else {
		fmt.Println("Connect database success!")
		fmt.Println(db)
	}
}

func OperationDatabase() {

}
