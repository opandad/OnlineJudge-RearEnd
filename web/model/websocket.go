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

	//分页查询
	Page Page `json:"page"`

	ProblemID int `json:"problemID"`
	ContestID int `json:"contestID"`

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
	Submits  []Submit   `json:"submit"`
}

type WebsocketUser struct {
	ID         int    `json:"userID"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

type Page struct {
	PageSize  int `json:"pageSize"`
	PageIndex int `json:"pageIndex"`
}
