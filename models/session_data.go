package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type SessionData struct {
	Email      string `form:"email"`
	UserID     int    `form:"userid"`
	Password   string `form:"password"`
	VerifyCode int    `form:"verifyCode"`
}
