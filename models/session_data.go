package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type SessionData struct {
	ID         int
	Email      string
	VerifyCode string
}
