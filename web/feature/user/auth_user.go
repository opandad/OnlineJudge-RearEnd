package user

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/web/model"
	"fmt"
)

func AuthUser(id int, authority *string) error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return err
	}
	var user struct {
		ID        int
		Authority string
	}
	err = mdb.Model(&model.User{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}
	authority = &user.Authority
	fmt.Println(user)
	return nil
}
