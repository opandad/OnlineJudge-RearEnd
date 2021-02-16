package contest

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/model"
	"errors"
	"fmt"
)

/*
正在开发中

@Title
contests.List

@description

@param

@return
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

	var contests []model.Contest
	err = mdb.Debug().Offset((websocketInputData.Page.PageIndex-1)*websocketInputData.Page.PageSize).Limit(websocketInputData.Page.PageSize).Select("id", "name", "start_time").Find(&contests).Error
	if err != nil {
		return err
	}
	websocketOutputData.Data.Contests = contests

	fmt.Println(contests)
	return nil
}
