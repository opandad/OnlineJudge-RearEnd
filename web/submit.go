package web

//非0验证
func (submit Submit) SubmitAnswer() HTTPStatus {
	if submit.ContestId != 0 {
		//查询是否有资格提交，没有资格返回错误
	}

	//执行判题机

	return HTTPStatus{}
}

func (submit Submit) SubmitQuery() ([]Submit, HTTPStatus) {
	return []Submit{}, HTTPStatus{}
}