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
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MYSQL_CONNECT *sql.DB = nil

func ReturnMysqlConfig() (string, string) {
	return "mysql", configs.DATABASE_MYSQL_USER + ":" + configs.DATABASE_MYSQL_PASSWORD + "@tcp(" + configs.DATABASE_MYSQL_SERVER_IP + ":" + configs.DATABASE_MYSQL_SERVER_PORT + ")/" + configs.DATABASE_MYSQL_NAME + "?charset=" + configs.DATABASE_MYSQL_CHARSET + "&parseTime=" + configs.DATABASE_MYSQL_PARSETIME + "&loc=" + configs.DATABASE_MYSQL_LOC
}

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
func InitMysqlDatabase() {
	var err error
	MYSQL_CONNECT, err = sql.Open(ReturnMysqlConfig())
	if err != nil {
		MYSQL_CONNECT.Close()
		log.Fatal("Connect to database fail:", err)
	}
	MYSQL_CONNECT.SetMaxIdleConns(configs.DATABASE_MYSQL_MAXIDLECONNS)
	MYSQL_CONNECT.SetMaxOpenConns(configs.DATABASE_MYSQL_MAXOPENCONNS)
	MYSQL_CONNECT.SetConnMaxLifetime(configs.DATABASE_MYSQL_CONNMAXLIFETIME)
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
	//debug模式
	if configs.DATABASE_LOG_MODE_DEBUG {
		// var err error
		// MYSQL_CONNECT, err = sql.Open(ReturnMysqlConfig())
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // 禁用彩色打印
			},
		)
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: MYSQL_CONNECT,
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
			Conn: MYSQL_CONNECT,
		}), &gorm.Config{})
		// if err != nil {
		// 	log.Fatal("Reconnect to database fail: ", err)
		// }
		// fmt.Println("Reconnect to database Success!")
		return gormDB, err
	}
}
