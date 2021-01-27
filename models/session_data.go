package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type SessionData struct {
	SessionID   string `json:"sessionID"`
	Account     string `json:"account"`
	UserID      int    `json:"userID"`
	Password    string `json:"password"`
	LoginByWhat string `json:"loginByWhat"`
	VerifyCode  string `json:"verifyCode"`
	Message     string `json:"msg"`
}
