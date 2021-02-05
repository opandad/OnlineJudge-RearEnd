package problem

/*
	@Title
	problem.list

	@Description
	题目细节功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| Detail              |    no    |    no	 |
*/

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/model"
)

/*
	正在开发中

	@Title
	problems.Detail

	@description
	显示题目细节返回给前端

	@param
	题目id

	@return
	题目详细信息 json格式
*/
func Detail(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return err
	}

	var problem model.Problem
	err = mdb.Debug().Where("id = ?", websocketInputData.ProblemID).First(&problem).Error
	if err != nil {
		return err
	}
	websocketOutputData.Problems = []model.Problem{problem}
	return nil
}
