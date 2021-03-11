package web

import (
	"OnlineJudge-RearEnd/api/database"
)

func (language Language) List() ([]Language, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Language{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "language.list",
			Method:      "",
		}
	}

	var languages []Language

	mdb.Find(&languages)

	return languages, HTTPStatus{
		IsError: false,
	}
}
