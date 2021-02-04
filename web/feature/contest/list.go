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
	var pageIndex, pageSize int
	pageIndex, pageSize = 1, 20
	if pageIndex <= 0 || pageSize <= 0 {
		return errors.New("非法输入")
	}

	var contests []model.Contest
	err = mdb.Debug().Offset((pageIndex-1)*pageSize).Limit(pageSize).Select("id", "name", "start_time").Find(&contests).Error
	if err != nil {
		return err
	}
	websocketOutputData.Contests = contests

	fmt.Println(contests)
	return nil
}
