package models

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type WebsocketInputData struct {
	websocketID string `json:"sessionID"`
	Account     string `json:"account"`
	UserID      int    `json:"userID"`
	Password    string `json:"password"`
	LoginByWhat string `json:"loginByWhat"`
	VerifyCode  string `json:"verifyCode"`
	Message     string `json:"msg"`
}

type WebsocketOutputData struct {
	UserID         int    `json:"userID"`
	WebsocketID    string `json:"sessionID"` //登录用
	Message        string `json:"msg"`       //About route
	HTTPStatusCode int    `json:"httpStatusCode"`
	Error          error  `json:"error"`
}
