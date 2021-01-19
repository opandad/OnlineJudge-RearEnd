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
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dsn := configs.DATABASE_USER + ":" + configs.DATABASE_PASSWORD + "@tcp(" + configs.DATABASE_SERVER_IP + ":" + configs.DATABASE_SERVER_PORT + ")/" + configs.DATABASE_NAME + "?charset=" + configs.DATABASE_CHARSET + "&parseTime=" + configs.DATABASE_PARSETIME + "&loc=" + configs.DATABASE_LOC
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
	db.SetMaxIdleConns(configs.DATABASE_MAXIDLECONNS)
	db.SetMaxOpenConns(configs.DATABASE_MAXOPENCONNS)
	db.SetConnMaxLifetime(configs.DATABASE_CONNMAXLIFETIME)
	fmt.Println("Init database success!")
}

/*
@Title
ReconnectDatabase

@description
返回mysql的*gorm.DB连接进行操作数据库，可以开启debug模式，在configs的database文件进行配置

@param
nil

@return
db (*gorm.DB)
*/
func ReconnectDatabase() *gorm.DB {
	mysqlDB := ConnectDatabase()

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
		if err != nil {
			log.Fatal("Reconnect to database fail: ", err)
		}
		fmt.Println("Reconnect to database success!")
		return gormDB.Session(&gorm.Session{Logger: newLogger})
	} else {
		gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: mysqlDB,
		}), &gorm.Config{})
		if err != nil {
			log.Fatal("Reconnect to database fail: ", err)
		}
		fmt.Println("Reconnect to database success!")
		return gormDB
	}
}
