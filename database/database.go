/*
	@Title
	database.go

	@Func List
	ConnectDatabase
	InitDatabase
	ReconnectDatabase
*/
package database

import (
	"OnlineJudge-RearEnd/configs"
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
@Title
ConnectDatabase

@description
返回一个Mysql数据库连接，这里可以接着修改成返回多种数据库配置

@param
nil

@return
db (*sql.DB)
*/
func ConnectDatabase() *sql.DB {
	driver := "mysql"
	dsn := configs.DATABASEUSER + ":" + configs.DATABASEPASSWORD + "@tcp(" + configs.DATABASEIP + ":" + configs.DATABASEPORT + ")/" + configs.DATABASENAME + "?charset=" + configs.DATABASECHARSET + "&parseTime=" + configs.DATABASEPARSETIME + "&loc=" + configs.DATABASELOC
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Connect to database fail:", err)
		return nil
	} else {
		log.Fatal("Connect to database Success!")
		return db
	}
}

/*
@Title
InitDatabase

@description
初始化连接数据库，可以通过configs包下的database.go文件进行连接池配置

@param
nil

@return
nil
*/
func InitDatabase() {
	db := ConnectDatabase()
	if db == nil {
		log.Fatal("Init database fail!")
		return
	}

	db.SetMaxIdleConns(configs.MAXIDLECONNS)
	db.SetMaxOpenConns(configs.MAXOPENCONNS)
	db.SetConnMaxLifetime(configs.CONNMAXLIFETIME)
}

/*
Error now，还未编写完成

@Title
ReconnectDatabase

@description
返回mysql的*gorm.DB连接进行操作数据库

@param
nil

@return
(*gorm.DB)
*/
func ReconnectDatabase() *gorm.DB {
	mysqlDB := ConnectDatabase()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Reconnect to database fail: ", err)
		return nil
	} else {
		log.Fatal("Reconnect to database success!")
		return gormDB
	}
}

/*
废弃代码
*/
// func ConnectDatabase() {
// 	dsn := configs.DATABASEUSER + ":" + configs.DATABASEPASSWORD + "@tcp(" + configs.DATABASEIP + ":" + configs.DATABASEPORT + ")/" + configs.DATABASENAME + "?charset=" + configs.DATABASECHARSET + "&parseTime=" + configs.DATABASEPARSETIME + "&loc=" + configs.DATABASELOC
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println("Connect database fail. Please check database connect config!")
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("Connect database success!")
// 		fmt.Println(db)
// 	}
// 	//设置连接池
// 	sqlDB, err := db.DB()
// 	sqlDB.SetMaxIdleConns(configs.MAXIDLECONNS)
// 	sqlDB.SetMaxOpenConns(configs.MAXOPENCONNS)
// 	sqlDB.SetConnMaxLifetime(configs.CONNMAXLIFETIME)
// }
