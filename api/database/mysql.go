/*
	@Title
	api/database/mysql.go

	@Description
	mysql数据库的连接管理以及连接池等功能

	@Func List

	| func name             | develop  | unit test |

	|----------------------------------------------|

	| ConnectMysqlDatabase  |    ok    |    ok	   |

	| InitMysqlDatabase     |    ok    |    ok	   |

	| ReconnectMysqlDatabase|    ok    |    ok	   |

	@config database path => ~/configs/mysql.go
*/
package database

import (
	"OnlineJudge-RearEnd/configs"
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
@Title
ConnectMysqlDatabase

@description
返回一个Mysql数据库连接，这里可以接着修改成返回多种数据库配置

@param
nil

@return
db (*sql.DB)
*/
func ConnectMysqlDatabase() *sql.DB {
	driver := "mysql"
	dsn := configs.DATABASE_MYSQL_USER + ":" + configs.DATABASE_MYSQL_PASSWORD + "@tcp(" + configs.DATABASE_MYSQL_SERVER_IP + ":" + configs.DATABASE_MYSQL_SERVER_PORT + ")/" + configs.DATABASE_MYSQL_NAME + "?charset=" + configs.DATABASE_MYSQL_CHARSET + "&parseTime=" + configs.DATABASE_MYSQL_PARSETIME + "&loc=" + configs.DATABASE_MYSQL_LOC
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Connect to database fail:", err)
	}
	return db
}

/*
@Title
InitMysqlDatabase

@description
初始化连接数据库，可以通过configs包下的database.go文件进行连接池配置

@param
nil

@return
nil
*/
func InitMysqlDatabase() {
	db := ConnectMysqlDatabase()
	db.SetMaxIdleConns(configs.DATABASE_MYSQL_MAXIDLECONNS)
	db.SetMaxOpenConns(configs.DATABASE_MYSQL_MAXOPENCONNS)
	db.SetConnMaxLifetime(configs.DATABASE_MYSQL_CONNMAXLIFETIME)
	fmt.Println("Init database success!")
}

/*
@Title
ReconnectMysqlDatabase

@description
返回mysql的*gorm.DB连接进行操作数据库，可以开启debug模式，在configs的database文件进行配置

@param
nil

@return
db (*gorm.DB)
*/
func ReconnectMysqlDatabase() (*gorm.DB, error) {
	mysqlDB := ConnectMysqlDatabase()

	//debug模式
	if configs.DATABASE_LOG_MODE_DEBUG {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // 禁用彩色打印
			},
		)
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: mysqlDB,
		}), &gorm.Config{
			Logger: newLogger,
		})
		// if err != nil {
		// 	log.Fatal("Reconnect to database fail: ", err)
		// }
		// fmt.Println("Reconnect to database Success!")
		return gormDB.Session(&gorm.Session{Logger: newLogger}), err
	} else {
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: mysqlDB,
		}), &gorm.Config{})
		// if err != nil {
		// 	log.Fatal("Reconnect to database fail: ", err)
		// }
		// fmt.Println("Reconnect to database Success!")
		return gormDB, err
	}
}
