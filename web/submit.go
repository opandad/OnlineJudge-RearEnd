package web

import "OnlineJudge-RearEnd/api/judger"

//非0验证
func (submit Submit) SubmitAnswer() HTTPStatus {

	//执行判题机
	judger.Judger("./data/problems/APlusB/problem.json", "./data/codes/APlusB/ac.c", "")

	return HTTPStatus{}
}

func (submit Submit) List() ([]Submit, HTTPStatus) {
	return []Submit{}, HTTPStatus{}
}
