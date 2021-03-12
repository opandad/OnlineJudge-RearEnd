package web

import (
	"OnlineJudge-RearEnd/api/judger"
	"OnlineJudge-RearEnd/configs"
	"fmt"
)

//非0验证
func (submit Submit) SubmitAnswer() HTTPStatus {

	//执行判题机
	result, err := judger.Judger(configs.JUDGER_WORK_PATH+"/problems/APlusB/problem.json", configs.JUDGER_WORK_PATH+"/codes/APlusB/ac.c", "")
	if err != nil {
		fmt.Println(err)

		return HTTPStatus{
			Message:     "服务器发生错误，请联系管理员处理",
			IsError:     true,
			SubMessage:  "判题发生错误",
			RequestPath: "submit.submit",
		}
	}

	fmt.Println(result)

	return HTTPStatus{
		Message:     "",
		IsError:     false,
		SubMessage:  "",
		RequestPath: "submit.submit",
	}
}

func (submit Submit) List() ([]Submit, HTTPStatus) {
	return []Submit{}, HTTPStatus{}
}
