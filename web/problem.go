package web

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/problem_data"
	"OnlineJudge-RearEnd/configs"
	"OnlineJudge-RearEnd/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

/*
	@Title
	problem

	@Description
	题目相关

	@Func List

	Class name: problem

	| func name           | develop | unit test |  bug  |

	|---------------------------------------------------|

	| insert单            |   yes   |    no	    |  no   |

	| insert多            |   no    |    no	    |  no   |

	| Delete单            |   yes   |    no	    |  no   |

	| Delete多            |   no    |    no	    |  no   |

	| Update              |   yes   |    no	    |  no   |

	| Detail              |   yes   |    no	    |  no   |

	| List                |   yes   |    no	    |  no   |
*/

/*
	bug list
	没有做权限管理
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
			Method:      "",
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
			Method:      "",
		}
	}

	problem_data.MoveUploadFile(problem.ID)

	return HTTPStatus{
		Message:     "题目添加成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.insert",
		Method:      "PostProblem",
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
			Method:      "",
		}
	}
	// var contestsHasProblem ContestsHasProblem
	// mdb.Where("problem_id = ?", problem.ID).Delete(&contestsHasProblem)
	err = mdb.Delete(&problem).Error
	if err != nil {
		fmt.Println(err)

		return HTTPStatus{
			Message:     "删除出错！",
			IsError:     true,
			ErrorCode:   403,
			SubMessage:  "problem delete error, error code is error",
			RequestPath: "problem.delete",
			Method:      "",
		}
	}

	return HTTPStatus{
		Message:     "删除成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.delete",
		Method:      "ReturnProblem",
	}
}

/*
	偶为修改前，奇数为修改后
	输入的为修改后的
*/
func (problem Problem) Update() HTTPStatus {
	if problem.ID <= 0 {
		return HTTPStatus{
			Message:     "输入的什么鬼东西",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.Update",
			Method:      "",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.Update",
			Method:      "",
		}
	}
	err = mdb.Save(&problem).Error
	if err != nil {
		return HTTPStatus{
			Message:     "更新失败",
			IsError:     true,
			ErrorCode:   1,
			SubMessage:  "update error, error code is error",
			RequestPath: "problem.update",
			Method:      "",
		}
	}

	return HTTPStatus{
		Message:     "更新成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.update",
		Method:      "PutProblem",
	}
}

/*
	查询可以优化

	bug
	查过头会报错

	题目，状态，总数
*/
func (problem Problem) List(pageIndex int, pageSize int) ([]Problem, HTTPStatus, int64) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Problem{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.list",
			Method:      "",
		}, 0
	}

	//分页查询
	if pageIndex <= 0 || pageSize <= 0 {
		return []Problem{}, HTTPStatus{
			Message:     "非法输入",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "page index or page size input error, error code is error",
			RequestPath: "problem.list",
			Method:      "",
		}, 0
	}

	var problems []Problem
	mdb.Debug().Offset((pageIndex-1)*pageSize).Limit(pageSize).Select("id", "name", "is_hide_to_user").Find(&problems)
	// if err != nil {
	// 	return []Problem{}, HTTPStatus{
	// 		Message:     "服务器出错啦，请稍后重新尝试。",
	// 		IsError:     true,
	// 		ErrorCode:   500,
	// 		SubMessage:  "query error",
	// 		RequestPath: "problem.list",
	// 		Method:      "",
	// 	}, 0
	// }
	var count int64
	mdb.Model(&Problem{}).Count(&count)

	return problems, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "GetProblemList",
	}, count
}

/*
	需要输入id
*/
func (problem Problem) Detail() (Problem, HTTPStatus) {
	if problem.ID <= 0 {
		return Problem{}, HTTPStatus{
			Message:     "输入的什么鬼东西",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.detail",
			Method:      "",
		}
	}

	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return Problem{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "problem.detail",
			Method:      "",
		}
	}

	if errors.Is(mdb.First(&problem).Error, gorm.ErrRecordNotFound) {
		return Problem{}, HTTPStatus{
			Message:     "没有这个题目",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "id error",
			RequestPath: "problem.detail",
			Method:      "",
		}
	}

	return problem, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "problem.detail",
		Method:      "GetProblemDetail",
	}
}

//<================= 其他函数 文件校验 ============>

func (problem Problem) CheckUploadFiles() bool {
	files, _ := ioutil.ReadDir(configs.JUDGER_UPLOAD_TEMP_FILE_PATH)

	num := len(files)

	if num%2 != 0 {
		os.RemoveAll(configs.JUDGER_UPLOAD_TEMP_FILE_PATH)
		fmt.Println("文件缺失")
		return false
	}

	type File struct {
		Name   string
		Suffix string
	}

	var file []File

	file = make([]File, num)

	for i, f := range files {
		strs := strings.Split(f.Name(), ".")
		file[i].Name = strs[0]
		file[i].Suffix = strs[1]
	}

	sort.SliceStable(file, func(i, j int) bool {
		if file[i].Suffix != file[j].Suffix {
			return file[i].Suffix < file[j].Suffix
		}

		return file[i].Name < file[j].Name
	})

	for i, j := 0, num/2; i < num/2; {
		if (file[i].Name == file[j].Name) && (file[i].Suffix == "in" && file[j].Suffix == "out") {
		} else {
			fmt.Println("文件格式错误")
			os.RemoveAll(configs.JUDGER_UPLOAD_TEMP_FILE_PATH)
			return false
		}

		i++
		j++
	}

	fmt.Println("校验文件成功")
	return true
}

func (problem Problem) MoveUploadFile(problemID int) {
	files, _ := ioutil.ReadDir(configs.JUDGER_UPLOAD_TEMP_FILE_PATH)

	pathExists, err := utils.PathExists(configs.JUDGER_WORK_PATH + strconv.Itoa(problemID) + "/")
	if err == nil && pathExists == false {
		os.Mkdir(configs.JUDGER_WORK_PATH+strconv.Itoa(problemID)+"/", os.ModePerm)
	} else if err != nil && pathExists == false {
		fmt.Println(err)
	}

	for _, f := range files {
		// fmt.Println(f.Name())
		os.Rename(configs.JUDGER_UPLOAD_TEMP_FILE_PATH+f.Name(), configs.JUDGER_WORK_PATH+strconv.Itoa(problemID)+"/"+f.Name())
	}
}

func (problem Problem) ReturnProblemDataConfig() []TestCase {
	var testCase []TestCase

	files, _ := ioutil.ReadDir(configs.JUDGER_UPLOAD_TEMP_FILE_PATH)

	var len int = len(files)

	testCase = make([]TestCase, len/2)

	i := 0
	for _, f := range files {
		strs := strings.Split(f.Name(), ".")
		if strs[1] == ".in" {
			testCase[i].Handle = strs[0]
			testCase[i].Name = "Test #" + strs[0]
			testCase[i].Input = strs[0] + ".in"
			testCase[i].Output = strs[0] + ".out"
			testCase[i].Enable = true
			i++
		}
	}

	return testCase
}
