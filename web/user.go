package web

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/verification"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

/*
	<==========================email账号相关==========================>
*/

/*
正在开发

@Title
Email.Login

@description
Email登录模块，通常用于email账号登录验证，返回用户权限给前端

@param
password, websocketID (string, string)

@return
user, error (User, error)
*/
func (account Email) Login(password string, websocketID string) (User, error) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return User{}, errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return User{}, errors.New("redis数据库连接失败！")
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//未查询到账号
	if errors.Is(tx.Where("email = ?", account.Email).Find(&account).Error, gorm.ErrRecordNotFound) {
		return User{}, gorm.ErrRecordNotFound
	}

	tx.Model(&account).Association("User").Find(&account.User)

	//TODO 加密，和数据库比较

	if password != account.User.Password {
		return User{}, gorm.ErrRecordNotFound
	}

	//TODO 记录进redis中，证明已经登录过
	var userData UserData
	userData.WebsocketID = websocketID

	err = rdb.Set(ctx, string(account.User.ID), userData, time.Minute*30).Err()
	if err != nil {
		return User{}, errors.New("redis数据库添加失败！")
	}

	return account.User, nil
}

/**/
func (account Email) Logout() error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return errors.New("redis数据库连接失败！")
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//未查询到账号
	if errors.Is(tx.Where("email = ?", account.Email).Find(&account).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}

	ctx = context.Background()
	rdb.Del(ctx, string(account.UserID))

	return nil
}

/**/
func (account Email) Regist(websocketID string, password string, verifiCode string) error {
	//检查数据库是否能够正常连接
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return errors.New("redis数据库连接失败！")
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return errors.New("mysql数据库连接失败！")
	}

	ctx := context.Background()
	//尝试取数据
	userDataJSON, err := rdb.Get(ctx, account.Email).Result()
	if err != nil {
		return errors.New("没有这个帐号的验证码")
	}

	//取数据转换json
	var userData UserData
	err = json.Unmarshal([]byte(userDataJSON), &userData)
	if err != nil {
		return errors.New("redis数据转换json出错")
	}

	//将取出来的验证码比较
	if userData.VerifyCode != verifiCode {
		return errors.New("验证码错误")
	}

	//加入数据到mysql
	account.User.Name = verification.RandVerificationCode()
	err = mdb.Create(&account).Error
	if err != nil {
		fmt.Println(err)
		return errors.New("添加用户失败")
	}

	//redis抹掉验证码
	rdb.Del(ctx, string(account.User.ID))

	//自动登录
	rdb.Set(ctx, string(account.User.ID), &UserData{
		WebsocketID: websocketID,
		Authority:   account.User.Authority,
	}, time.Minute*30)

	return nil
}

/**/
func (account Email) AuthLogin() error {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return errors.New("redis数据库连接失败！")
	}
	ctx := context.Background()
	tx := mdb.WithContext(ctx)
	//未查询到账号
	if errors.Is(tx.Where("email = ?", account.Email).Find(&account).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	//找rdb授权
	ctx = context.Background()
	res, err := rdb.Get(ctx, string(account.UserID)).Result()
	if err != nil {
		return err
	}
	return nil
}

/*
	<==========================user账号相关=============================>
*/

/*
正在开发

@Title
User.Login

@description
User登录模块，通常用于自动登录验证，返回用户权限给前端

@param
ID, password (int, string)

@return
authority, error (string, error)
*/
func (account User) Login(websocketID string) (User, error) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return User{}, errors.New("mysql数据库连接失败，请重新检查mysql数据库配置！")
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return User{}, errors.New("redis数据库连接失败！")
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//TODO 加密，和数据库比较

	//未查询到账号
	if errors.Is(tx.Where("email = ? AND password = ?", account.ID, account.Password).Find(&account).Error, gorm.ErrRecordNotFound) {
		return User{}, gorm.ErrRecordNotFound
	}

	//TODO 记录进redis中，证明已经登录过
	var userData UserData
	userData.WebsocketID = websocketID

	err = rdb.Set(ctx, string(account.ID), userData, time.Minute*30).Err()
	if err != nil {
		return User{}, errors.New("redis数据库添加失败！")
	}

	return account, nil
}

/**/
func (account User) Logout() error {
	return nil
}

/**/
func (account User) Regist() error {
	return nil
}

/**/
func (account User) AuthLogin() error {
	return nil
}

/**/
func (account User) Insert() error {
	return nil
}

/**/
func (account User) Update() error {
	return nil
}
