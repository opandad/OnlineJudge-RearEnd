package database

import (
	"OnlineJudge-RearEnd/config"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type JSON json.RawMessage

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Submit struct {
	id            uint
	submit_state  string
	language      string
	run_time      time.Time
	submit_time   time.Time
	problems_id   int
	contest_id    int
	language_name string
	user_id       int
}

type User struct {
	id        uint   `gorm:"column:id";primary_key`
	email     string `gorm:"column:email"`
	name      string `gorm:"column:name"`
	password  string `gorm:"column:password"`
	authority string `gorm:"column:authority"`
	user_info JSON   `gorm:"column:user_info"`
}

var db *gorm.DB

func ConnectDatabase() {
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

	//sql查询测试
	var submit []User
	db.Table("user").Find(&submit)
	fmt.Println(submit)
}

func GetDatabaseConnection() *gorm.DB {
	return db
}
