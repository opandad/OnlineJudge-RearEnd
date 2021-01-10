package configs

import (
	"time"
)

//Database base config
const DATABASEIP string = "192.168.121.131"
const DATABASEPORT string = "3306"
const DATABASEUSER string = "online_judge_admin"
const DATABASEPASSWORD string = "qweasd"
const DATABASENAME string = "online_judge"
const DATABASECHARSET string = "utf8mb4"
const DATABASEPARSETIME string = "true"
const DATABASELOC string = "Local"

//Database connection pool config
const MAXIDLECONNS int = 10
const MAXOPENCONNS int = 100
const CONNMAXLIFETIME time.Duration = time.Hour
