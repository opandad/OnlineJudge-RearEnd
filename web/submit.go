package web

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/judger"
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/configs"
	"OnlineJudge-RearEnd/utils"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"gorm.io/gorm"
)

//非0验证
func (submit Submit) SubmitAnswer() HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "submit.submit",
			Method:      "",
		}
	}

	var count int64
	var user []User
	var problem []Problem
	var contest Contest
	if submit.ContestId != 0 {
		err = mdb.Where("id = ?", submit.ContestId).First(&contest).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return HTTPStatus{
				Message:     "比赛不存在",
				IsError:     true,
				ErrorCode:   500,
				SubMessage:  "比赛不存在",
				RequestPath: "submit.submit",
				Method:      "",
			}
		}
		mdb.Model(&contest).Where("problem_id = ?", submit.ProblemId).Association("Problems").Find(&problem)
		if len(problem) <= 0 {
			return HTTPStatus{
				Message:     "题目不存在",
				IsError:     true,
				ErrorCode:   500,
				SubMessage:  "题目不存在",
				RequestPath: "submit.submit",
				Method:      "",
			}
		}
		mdb.Model(&contest).Where("user_id = ?", submit.UserId).Association("Users").Find(&user)
		if len(user) <= 0 {
			return HTTPStatus{
				Message:     "用户没有参加比赛",
				IsError:     true,
				ErrorCode:   500,
				SubMessage:  "用户没有参加比赛",
				RequestPath: "submit.submit",
				Method:      "",
			}
		}
	} else {
		mdb.Where("id = ?", submit.ProblemId).First(&problem).Count(&count)
		if count <= 0 {
			return HTTPStatus{
				Message:     "题目不存在",
				IsError:     true,
				ErrorCode:   500,
				SubMessage:  "题目不存在",
				RequestPath: "submit.submit",
				Method:      "",
			}
		}
	}

	var language Language
	mdb.Where("id = ?", submit.LanguageId).Find(&language).Count(&count)
	if count <= 0 {
		return HTTPStatus{
			Message:     "语言不存在",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "语言不存在",
			RequestPath: "submit.submit",
			Method:      "",
		}
	}

	//TODO 将code写入文件中
	var submitCodeFileName = verification.RandVerificationCode() + verification.Snowflake() + verification.RandVerificationCode()

	f, err := os.Create(configs.JUDGER_SUBMIT_PATH + submitCodeFileName) //创建文件
	if err != nil {
		fmt.Println("file create fail")
		return HTTPStatus{
			Message:     "服务器内部错误，请稍后尝试",
			IsError:     true,
			SubMessage:  "file create error",
			RequestPath: "submit.submitAnswer",
		}
	}
	//将文件写进去
	n, err := io.WriteString(f, submit.SubmitCode)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器内部错误，请稍后尝试",
			IsError:     true,
			SubMessage:  "file create error",
			RequestPath: "submit.submitAnswer",
		}
	}
	fmt.Println("写入了n个字节: ", n)
	defer f.Close()

	//执行判题机
	result, err := judger.Judger(configs.JUDGER_WORK_PATH+problem[0].JudgeerInfo.ProblemPath, configs.JUDGER_SUBMIT_PATH+submitCodeFileName, language.Language)
	if err != nil {
		fmt.Println(err)

		return HTTPStatus{
			Message:     "服务器发生错误，请联系管理员处理",
			IsError:     true,
			SubMessage:  "判题发生错误",
			RequestPath: "submit.submit",
		}
	}

	submit.SubmitState = result
	submit.SubmitTime = time.Now().Format(utils.StandardTimeFormat)
	submit.SubmitInfo.CodeFileName = submitCodeFileName

	if !(result == "Accepted" || result == "Pending") {
		//error
		submit.IsError = true
	} else {
		//not error
		submit.IsError = false
	}

	mdb.Save(&submit)

	return HTTPStatus{
		Message:     "提交成功",
		IsError:     false,
		SubMessage:  "",
		RequestPath: "submit.submit",
	}
}

func (submit Submit) List(pageIndex int, pageSize int) ([]Submit, HTTPStatus, int64) {
	if pageIndex <= 0 || pageSize <= 0 {
		return []Submit{}, HTTPStatus{
			Message:     "输入错误",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "pageIndex or pageSize error",
			RequestPath: "submit.list",
			Method:      "",
		}, 0
	}

	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Submit{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "submit.list",
			Method:      "",
		}, 0
	}

	var submits []Submit
	var total int64
	mdb.Table("submits").Where(&submit).Count(&total).Order("id desc").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&submits)

	return submits, HTTPStatus{
		Message:     "",
		IsError:     false,
		RequestPath: "submit.list",
	}, total
}
