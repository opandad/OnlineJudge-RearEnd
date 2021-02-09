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
	if websocketInputData.Page.PageIndex <= 0 || websocketInputData.Page.PageSize <= 0 {
		return errors.New("非法输入")
	}

	var problems []model.Problem
	err = mdb.Debug().Offset((websocketInputData.Page.PageIndex-1)*websocketInputData.Page.PageSize).Limit(websocketInputData.Page.PageSize).Select("id", "name", "accept", "fail").Find(&problems).Error
	if err != nil {
		return err
	}
	websocketOutputData.Data.Problems = problems

	// fmt.Println(problems)
	return nil
}
