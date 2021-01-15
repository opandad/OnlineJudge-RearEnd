/*
	@Title
	api/database.go

	@Description
	数据库的连接管理以及连接池等功能

	@Func List

	| func name         | develop  | unit test |

	|------------------------------------------|

	| ConnectDatabase   |    ok    |    ok	   |

	| InitDatabase      |    ok    |    ok	   |

	| ReconnectDatabase |    ok    |    ok	   |

	@config database path => ~/configs/database.go
*/
package database

import (
	"OnlineJudge-RearEnd/configs"
	"database/sql"
	"fmt"
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
	}
	fmt.Println("Connect to database Success!")
	return db
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
	db.SetMaxIdleConns(configs.MAXIDLECONNS)
	db.SetMaxOpenConns(configs.MAXOPENCONNS)
	db.SetConnMaxLifetime(configs.CONNMAXLIFETIME)
	fmt.Println("Init database success!")
}

/*
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
	}

	fmt.Println("Reconnect to database success!")
	return gormDB
}
