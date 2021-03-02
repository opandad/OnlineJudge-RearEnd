package web

import "OnlineJudge-RearEnd/api/database"

/*
	@Title
	problem

	@Description
	题目相关

	@Func List

	Class name: problem

	| func name           | develop | unit test |  bug  |

	|---------------------------------------------------|

	| insert单            |   no    |    no	    |  no   |

	| insert多            |   no    |    no	    |  no   |

	| Delete单            |   no    |    no	    |  no   |

	| Delete多            |   no    |    no	    |  no   |

	| Update              |   no    |    no	    |  no   |

	| Query               |   no    |    no	    |  no   |
*/

func (problem Problem) Insert() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.insert",
			Method:      "post",
		}
	}
	err = mdb.Create(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql insert fail",
			RequestPath: "problem.insert",
			Method:      "post",
		}
	}

	return HTTPStatus{
		Message:     "题目添加成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.insert",
		Method:      "post",
	}
}

/*
	@input
	problem.ID

	bug list
	可能会触发批量delete
*/
func (problem Problem) Delete() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}
	err = mdb.Delete(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "删除出错！",
			IsError:     true,
			ErrorCode:   403,
			SubMessage:  "problem delete error, error code is error",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}

	return HTTPStatus{
		Message:     "删除成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.delete",
		Method:      "delete",
	}
}

/*
	偶为修改前，奇数为修改后
	输入的为修改后的
*/
func (problem Problem) Update(updateProblem Problem) HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.delete",
			Method:      "delete",
		}
	}
	mdb.First(&problem)

	return HTTPStatus{}
}

/*
	查询可以优化
*/
func (problem Problem) QueryIndex(pageIndex int, pageSize int) ([]Problem, HTTPStatus) {
	return []Problem{}, HTTPStatus{}
}

func (problem Problem) QueryDetail() (Problem, HTTPStatus) {
	return Problem{}, HTTPStatus{}
}
