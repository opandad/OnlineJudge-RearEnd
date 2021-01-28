package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type WebsocketInputData struct {
	SessionID   int64  `json:"sessionID"`
	Account     string `json:"account"`
	UserID      int    `json:"userID"`
	Password    string `json:"password"`
	LoginByWhat string `json:"loginByWhat"`
	VerifyCode  string `json:"verifyCode"`
	Message     string `json:"msg"`
}

type WebsocketOutputData struct {
	UserID    int    `json:"userID"`
	SessionID int64  `json:"sessionID"` //登录用
	Message   string `json:"msg"`       //About route
	Error     error  `json:"error"`
}
