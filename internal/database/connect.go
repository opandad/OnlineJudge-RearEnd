/*
	功能：返回一个数据库连接
*/
package database

import (
	"OnlineJudge-RearEnd/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() {
	dsn := configs.DATABASEUSER + ":" + configs.DATABASEPASSWORD + "@tcp(" + configs.DATABASEIP + ":" + configs.DATABASEPORT + ")/" + configs.DATABASENAME + "?charset=" + configs.DATABASECHARSET + "&parseTime=" + configs.DATABASEPARSETIME + "&loc=" + configs.DATABASELOC

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
	sqlDB.SetMaxIdleConns(configs.MAXIDLECONNS)
	sqlDB.SetMaxOpenConns(configs.MAXOPENCONNS)
	sqlDB.SetConnMaxLifetime(configs.CONNMAXLIFETIME)

	//sql查询测试
	/* User查询测试 */
	// var submit []models.User
	// db.Debug().Find(&submit)
	// fmt.Println("id:", submit[0].ID, "email:", submit[0].Email, "name:", submit[0].Name, "password:", submit[0].Password, "authority:", submit[0].Authority, "userinfo:", submit[0].UserInfo)
	// fmt.Println("id:", submit[1].ID, "email:", submit[1].Email, "name:", submit[1].Name, "password:", submit[1].Password, "authority:", submit[1].Authority, "userinfo:", submit[1].UserInfo)
}

func GetDatabaseConnection() *gorm.DB {
	return db
}
