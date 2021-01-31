package model

/*
	db:0
	负责存储验证码信息
	key: account
	value: struct
*/
type UserOnlineData struct {
	WebsocketID string `json:"websocketID"`
	VerifyCode  string `json:"verifyCode"`
}
