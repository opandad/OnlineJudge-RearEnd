package web

import (
	"OnlineJudge-RearEnd/api/database"
	"context"
	"fmt"
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

	105,153
	213,238
*/

/*
	bug
	查过头会报错
*/
func (contest Contest) List(pageIndex int, pageSize int) ([]Contest, HTTPStatus, int) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Contest{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.list",
			Method:      "",
		}, 0
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
		}, 0
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
		}, 0
	}

	return contests, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "GetContestList",
	}, len(contests)
}

/*
	input contest.id
*/
func (contest Contest) Detail(userID int) (Contest, []Problem, []Language, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return Contest{}, []Problem{}, []Language{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.detail",
			Method:      "",
		}
	}

	//链式操作
	var users []User
	ctx := context.Background()
	tx := mdb.WithContext(ctx)
	tx.Model(&contest).Where("user_id = ?", userID).Association("Users").Find(&users)

	if len(users) <= 0 {
		return Contest{}, []Problem{}, []Language{}, HTTPStatus{
			Message:     "此用户没有参加比赛",
			IsError:     true,
			SubMessage:  "此用户没有参加比赛",
			RequestPath: "contest.detail",
		}
	}

	var problems []Problem
	var languages []Language

	tx.Model(&contest).Association("Problems").Find(&problems)
	tx.Model(&contest).Association("Languages").Find(&languages)

	// fmt.Println("problems: ", problems)
	// fmt.Println("languages", languages)

	return contest, problems, languages, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "contest.detail",
		Method:      "GetContestDetail",
	}
}

// /**/
// func (contest Contest) Rank() ([]Submit, HTTPStatus) {
// 	// mdb, err := database.ReconnectMysqlDatabase()
// 	// if err != nil {
// 	// 	return []Submit{}, HTTPStatus{
// 	// 		Message:     "服务器出错啦，请稍后重新尝试。",
// 	// 		IsError:     true,
// 	// 		ErrorCode:   500,
// 	// 		SubMessage:  "mysql database connect fail",
// 	// 		RequestPath: "contest.rank",
// 	// 		Method:      "",
// 	// 	}
// 	// }

// 	return []Submit{}, HTTPStatus{}
// }

/*
	取contest当中的announcement
*/
// func (contest Contest) Notice() HTTPStatus {
// 	return HTTPStatus{}
// }

func (contest Contest) Insert(problems []Problem, languages []Language, users []User) HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.insert",
			Method:      "",
		}
	}

	fmt.Println("contest insert")
	fmt.Println(contest)
	fmt.Println(problems)
	fmt.Println(languages)
	fmt.Println(users)

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	tx.Table("contests").Create(&contest)
	tx.Model(&contest).Association("Languages").Append(&languages)
	tx.Model(&contest).Association("Problems").Append(&problems)
	tx.Model(&contest).Association("Users").Append(&users)

	return HTTPStatus{
		Message:     "添加成功",
		IsError:     false,
		RequestPath: "contest insert",
	}
}

func (contest Contest) Update(problems []Problem, languages []Language, users []User) HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.update",
			Method:      "",
		}
	}

	fmt.Println("contest update")
	fmt.Println(contest)
	fmt.Println(problems)
	fmt.Println(languages)
	fmt.Println(users)

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	tx.Table("contests").Save(&contest)
	tx.Model(&contest).Association("Languages").Replace(&languages)
	tx.Model(&contest).Association("Problems").Replace(&problems)
	tx.Model(&contest).Association("Users").Replace(&users)

	return HTTPStatus{
		Message:     "修改成功",
		IsError:     false,
		RequestPath: "contest update",
	}
}

/*
	input:id

	可能有bug
*/
// func (contest Contest) Delete() HTTPStatus {
// 	mdb, err := database.ReconnectMysqlDatabase()
// 	if err != nil {
// 		return HTTPStatus{
// 			Message:     "服务器出错啦，请稍后重新尝试。",
// 			IsError:     true,
// 			ErrorCode:   500,
// 			SubMessage:  "mysql database connect fail",
// 			RequestPath: "contest.rank",
// 			Method:      "",
// 		}
// 	}

// 	mdb.Model(&UsersJoinContest{}).Where("contests_id = ?", contest.ID).Delete(&UsersJoinContest{})
// 	mdb.Model(&ContestsSupportLanguage{}).Where("contests_id = ?", contest.ID).Delete(&ContestsSupportLanguage{})
// 	mdb.Model(&ContestsHasProblem{}).Where("contests_id = ?", contest.ID).Delete(&ContestsHasProblem{})

// 	mdb.Delete(&contest)

// 	return HTTPStatus{}
// }

func (contest Contest) GetEdit() (Contest, []Problem, []Language, HTTPStatus, []User, []Language) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return Contest{}, []Problem{}, []Language{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "contest.detail",
			Method:      "",
		}, []User{}, []Language{}
	}

	//链式操作
	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	var users []User
	var problems []Problem
	var languages []Language
	var selectLanguages []Language

	tx.Where(&contest).First(&contest)
	tx.Model(&contest).Select("id").Association("Users").Find(&users)
	tx.Model(&contest).Select("id").Association("Problems").Find(&problems)
	tx.Model(&contest).Association("Languages").Find(&languages)
	tx.Find(&selectLanguages)

	// fmt.Println("problems: ", problems)
	// fmt.Println("languages", languages)

	return contest, problems, languages, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "contest.detail",
		Method:      "GetContestDetail",
	}, users, selectLanguages
}
