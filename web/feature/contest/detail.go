package contest

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/model"
)

/*
正在开发中

@Title
contests.Detail

@description

@param

@return
*/
func Detail(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return err
	}

	var contest model.Contest

	err = mdb.Debug().Where("id = ?", websocketInputData.ProblemID).First(&contest).Error
	if err != nil {
		return err
	}
	websocketOutputData.Data.Contests = []model.Contest{contest}
	return nil
}
