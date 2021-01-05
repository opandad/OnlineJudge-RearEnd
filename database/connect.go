package database

import (
	"OnlineJudge-RearEnd/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type DatabaseConfig struct {
// 	ServerIP       string `json:"ServerIP"`
// 	Port           string `json:"ServerPort"`
// 	User           string `json:"User"`
// 	Password       string `json:"Password"`
// 	Database       string `json:"Database"`
// 	DatabaseEncode string `json:"DatabaseEncode"`
// }

var db *gorm.DB

func ConnectDatabase() {
	// file, err := os.Open("./config/database.json")
	// if err != nil {
	// 	fmt.Println("Open database config file fail!", err.Error())
	// 	return
	// } else {
	// 	fmt.Println("Open database config file success!")
	// }
	// defer file.Close()

	// var dbConfig DatabaseConfig
	// byteValue, _ := ioutil.ReadAll(file)
	// json.Unmarshal([]byte(byteValue), &dbConfig)

	// dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.ServerIP + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=" + dbConfig.DatabaseEncode + "&parseTime=True&loc=Local"
	dsn := config.DATABASEUSER + ":" + config.DATABASEPASSWORD + "@tcp(" + config.DATABASEIP + ":" + config.DATABASEPORT + ")/" + config.DATABASENAME + "?charset=" + config.DATABASECHARSET + "&parseTime=" + config.DATABASEPARSETIME + "&loc=" + config.DATABASELOC
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect database fail. Please check database connect config!")
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connect database success!")
		fmt.Println(db)
	}

	//设置连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(config.MAXIDLECONNS)
	sqlDB.SetMaxOpenConns(config.MAXOPENCONNS)
	sqlDB.SetConnMaxLifetime(config.CONNMAXLIFETIME)
}

func GetDatabaseConnection() *gorm.DB {
	return db
}
