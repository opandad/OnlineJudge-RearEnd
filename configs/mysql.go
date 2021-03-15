package configs

import (
	"time"
)

const (
	//Database base config
	DATABASE_MYSQL_SERVER_IP   string = "127.0.0.1"
	DATABASE_MYSQL_SERVER_PORT string = "3306"
	DATABASE_MYSQL_USER        string = "online_judge_admin"
	DATABASE_MYSQL_PASSWORD    string = "qweasd"
	DATABASE_MYSQL_NAME        string = "online_judge"
	DATABASE_MYSQL_CHARSET     string = "utf8mb4"
	DATABASE_MYSQL_PARSETIME   string = "false"
	DATABASE_MYSQL_LOC         string = "Local"

	//Database connection pool config
	DATABASE_MYSQL_MAXIDLECONNS    int           = 10
	DATABASE_MYSQL_MAXOPENCONNS    int           = 100
	DATABASE_MYSQL_CONNMAXLIFETIME time.Duration = time.Hour

	//Database log mode debug
	DATABASE_LOG_MODE_DEBUG bool = true
)
