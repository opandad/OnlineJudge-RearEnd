package web

import (
	"OnlineJudge-RearEnd/api/database"
	"context"
	"errors"

	"gorm.io/gorm"
)

/*
	要求
	1：过滤用户
	2：排名系统
	3：题目列表系统
	4：正确率系统
	5：增删改查比赛
*/

/*
	@Title
	contest

	@Description
	比赛相关

	@Func List

	Class name: contest

	| func name           | develop | unit test |

	|-------------------------------------------|

	| List                |   yes   |    no     |

	| Detail              |   yes   |    no     |

	| Rank                |   no    |    no     |

	| Notice              |   no    |    no     |

	| Insert              |   no    |    no     |

	| Update              |   no    |    no     |

	| Delete              |   yes   |    no     |
*/

/*
	bug
	查过头会报错
*/
func (contest Contest) List(pageIndex int, pageSize int) ([]Contest, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Contest{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.list",
			Method:      "",
		}
	}

	//分页查询
	if pageIndex <= 0 || pageSize <= 0 {
		return []Contest{}, HTTPStatus{
			Message:     "非法输入",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "page index or page size input error, error code is error",
			RequestPath: "contest.list",
			Method:      "",
		}
	}

	var contests []Contest
	err = mdb.Debug().Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&contests).Error
	if err != nil {
		return []Contest{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "query error",
			RequestPath: "contest.list",
			Method:      "",
		}
	}

	return contests, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "GetContestList",
	}
}

/**/
func (contest Contest) Detail(userID int) (Contest, []ContestsHasProblem, []ContestsSupportLanguage, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return Contest{}, []ContestsHasProblem{}, []ContestsSupportLanguage{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.list",
			Method:      "",
		}
	}

	//链式操作
	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//判断是否参加比赛
	var usersJoinContest UsersJoinContest
	err = tx.Where("users_id = ? AND contests_id = ?", userID, contest.ID).Find(&usersJoinContest).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Contest{}, []ContestsHasProblem{}, []ContestsSupportLanguage{}, HTTPStatus{
			Message:     "没有参加比赛",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "没有找到用户参加比赛的数据，error code is error",
			RequestPath: "contest.detail",
			Method:      "",
		}
	}

	tx.Where("id = ?", contest.ID).Find(&contest)

	var contestsHasProblems []ContestsHasProblem
	tx.Where("contests_id = ?", contest.ID).Find(&contestsHasProblems)

	var contestsSupportLanguages []ContestsSupportLanguage
	tx.Where("contests_id = ?", contest.ID).Find(&contestsSupportLanguages)

	return contest, contestsHasProblems, contestsSupportLanguages, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "contest.detail",
		Method:      "GetContestDetail",
	}
}

/**/
func (contest Contest) Rank() ([]Submit, HTTPStatus) {
	// mdb, err := database.ReconnectMysqlDatabase()
	// if err != nil {
	// 	return []Submit{}, HTTPStatus{
	// 		Message:     "服务器出错啦，请稍后重新尝试。",
	// 		IsError:     true,
	// 		ErrorCode:   500,
	// 		SubMessage:  "mysql database connect fail",
	// 		RequestPath: "contest.rank",
	// 		Method:      "",
	// 	}
	// }

	return []Submit{}, HTTPStatus{}
}

/*
	取contest当中的announcement
*/
// func (contest Contest) Notice() HTTPStatus {
// 	return HTTPStatus{}
// }

func (contest Contest) Insert() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.rank",
			Method:      "",
		}
	}

	mdb.Create(&contest)

	return HTTPStatus{}
}

func (contest Contest) Update() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.rank",
			Method:      "",
		}
	}

	mdb.Save(&contest)

	return HTTPStatus{}
}

/*
	input:id

	可能有bug
*/
func (contest Contest) Delete() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.rank",
			Method:      "",
		}
	}

	mdb.Model(&UsersJoinContest{}).Where("contests_id = ?", contest.ID).Delete(&UsersJoinContest{})
	mdb.Model(&ContestsSupportLanguage{}).Where("contests_id = ?", contest.ID).Delete(&ContestsSupportLanguage{})
	mdb.Model(&ContestsHasProblem{}).Where("contests_id = ?", contest.ID).Delete(&ContestsHasProblem{})

	mdb.Delete(&contest)

	return HTTPStatus{}
}
