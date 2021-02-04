package model

/*
	session是存放在redis当中的
	["用户id"]:["用户数据"]
*/
type WebsocketInputData struct {
	//web
	WebsocketID string `json:"websocketID"`
	RequestPath string `json:"requestPath"`

	//user
	User WebsocketUser `json:"user"`

	Submit Submit `json:"submit"`
}

type WebsocketOutputData struct {
	//web
	WebsocketID    string `json:"websocketID"` //登录用
	Message        string `json:"msg"`         //About route
	HTTPStatusCode int    `json:"httpStatusCode"`
	IsError        bool   `json:"isError"`

	//user
	User User `json:"user"`

	//about feature
	Problems []Problem  `json:"problem"`
	Contests []Contest  `json:"contest"`
	Language []Language `json:"language"`
}

type WebsocketUser struct {
	Account     string `json:"account"`
	Password    string `json:"password"`
	LoginByWhat string `json:"loginByWhat"`
	VerifyCode  string `json:"verifyCode"`
}
