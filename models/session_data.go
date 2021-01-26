package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type SessionData struct {
	Email      string `form:"email" json:"email"`
	UserID     int    `form:"userID" json:"userID"`
	Password   string `form:"password" json:"password"`
	VerifyCode string `form:"verifyCode" json:"verifyCode"`
}
