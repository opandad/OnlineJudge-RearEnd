package configs

import "time"

const (
	REAREND_SERVER_IP   string = "127.0.0.1"
	REAREND_SERVER_PORT string = "8080"

	/*
		websocket
		关于验证码有效时间以及用户登录有效时间
	*/
	VERIFYCODE_LIFT_TIME time.Duration = time.Minute * 10
	USER_LOGIN_LIFT_TIME time.Duration = time.Minute * 60
)
