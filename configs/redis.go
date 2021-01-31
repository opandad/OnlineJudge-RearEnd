package configs

import "time"

const (
	//redis database config
	DATABASE_REDIS_SERVER_IP   string = "127.0.0.1"
	DATABASE_REDIS_SERVER_PORT string = "6379"
	DATABASE_REDIS_PASSWORD    string = "qweasd"

	//redis data life time
	DATABASE_REDIS_DATA_LIFE_TIME time.Duration = time.Hour
)
