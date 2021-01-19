package configs

import (
	"time"
)

const (
	//Database base config
	DATABASE_SERVER_IP   string = "192.168.121.131"
	DATABASE_SERVER_PORT string = "3306"
	DATABASE_USER        string = "online_judge_admin"
	DATABASE_PASSWORD    string = "qweasd"
	DATABASE_NAME        string = "online_judge"
	DATABASE_CHARSET     string = "utf8mb4"
	DATABASE_PARSETIME   string = "true"
	DATABASE_LOC         string = "Local"

	//Database connection pool config
	DATABASE_MAXIDLECONNS    int           = 10
	DATABASE_MAXOPENCONNS    int           = 100
	DATABASE_CONNMAXLIFETIME time.Duration = time.Hour

	//Database log mode debug
	DATABASE_LOG_MODE_DEBUG bool = true
)
