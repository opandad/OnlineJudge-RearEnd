package problem

/*
	@Title
	problem.list

	@Description
	显示题目功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| List                |    ok    |    no	 |
*/

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/model"
	"errors"
)

/*
需要进行测试

@Title
problems.List

@description
problem请求，返回分页题目数据

@param
pageIndex, pageSize(int, int)

@return
problem list([]model.Problems)
*/
func List(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return err
	}

	//分页查询
	var pageIndex, pageSize int
	pageIndex, pageSize = 1, 20
	if pageIndex <= 0 || pageSize <= 0 {
		return errors.New("非法输入")
	}

	var problems []model.Problem
	err = mdb.Debug().Offset((pageIndex-1)*pageSize).Limit(pageSize).Select("id", "name", "accept", "fail").Find(&problems).Error
	if err != nil {
		return err
	}
	websocketOutputData.Problems = problems

	// fmt.Println(problems)
	return nil
}
