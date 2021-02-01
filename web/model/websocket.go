package model

/*
	session是存放在redis当中的
	["机子id"]:["用户数据"]
*/
type WebsocketInputData struct {
	//web
	WebsocketID string `json:"websocketID"`
	Message     string `json:"msg"`

	//user
	Account     string `json:"account"`
	User        User   `json:"user"`
	LoginByWhat string `json:"loginByWhat"`
	VerifyCode  string `json:"verifyCode"`

	Submit Submit `json:"submit"`
}

type WebsocketOutputData struct {
	//web
	WebsocketID    string `json:"websocketID"` //登录用
	Message        string `json:"msg"`         //About route
	HTTPStatusCode int    `json:"httpStatusCode"`

	//user
	User User `json:"user"`

	//about feature
	Problems []Problem  `json:"problem"`
	Contests []Contest  `json:"contest"`
	Language []Language `json:"language"`
}
