package configs

import (
	"time"
)

const (
	//Database base config
	DATABASEIP        string = "192.168.121.131"
	DATABASEPORT      string = "3306"
	DATABASEUSER      string = "online_judge_admin"
	DATABASEPASSWORD  string = "qweasd"
	DATABASENAME      string = "online_judge"
	DATABASECHARSET   string = "utf8mb4"
	DATABASEPARSETIME string = "true"
	DATABASELOC       string = "Local"

	//Database connection pool config
	MAXIDLECONNS    int           = 10
	MAXOPENCONNS    int           = 100
	CONNMAXLIFETIME time.Duration = time.Hour
)
