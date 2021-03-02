package web

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/verification"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

/*
	@Title
	account interface

	@Description
	用户相关接口规范

	@Func List

	Class name: email

	| func name           | develop | unit test |  bug  |

	|---------------------------------------------------|

	| Login               |   yes   |    no	    |  yes  |

	| Logout              |   yes   |    no	    |  yes  |

	| Regist              |   yes   |    no	    |  no   |

	| AuthLogin           |   yes   |    no	    |  yes  |

	| SendVerifyCode      |   yes   |    no	    |  no   |

	Class name: user

	| func name           | develop  | unit test |

	|--------------------------------------------|

	| Login               |    yes   |    no	 |

	| Logout              |    yes   |    no	 |

	| AuthLogin           |    no    |    no	 |
*/
type account interface {
	Login(websocketID string) (int, HTTPStatus) //返回userID和HTTPStatus
	Logout(websocketID string) HTTPStatus
	//Regist(websocketID string, verifiCode string) (int, HTTPStatus) //返回userID和HTTPStatus
	AuthLogin(websocketID string) HTTPStatus
}

/*
	<==========================email账号相关==========================>
*/

/*
@Title
Email.Login

@description
Email登录模块，通常用于email账号登录验证，返回用户权限给前端

@param
password, websocketID (string, string)

@return
UserID, HTTPStatus
*/
func (account Email) Login(websocketID string) (int, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.login",
			Method:      "get",
		}
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.login",
			Method:      "get",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	var emailAccount Email
	emailAccount.Email = account.Email
	//未查询到账号
	if errors.Is(tx.Where("email = ?", emailAccount.Email).Find(&emailAccount).Error, gorm.ErrRecordNotFound) {
		return -1, HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "email.login",
			Method:      "get",
		}
	}

	tx.Model(&emailAccount).Association("User").Find(&emailAccount.User)

	//TODO 加密，和数据库比较

	if emailAccount.User.Password != account.User.Password {
		return -1, HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "email.login",
			Method:      "get",
		}
	}

	//数据记录并转换json
	var userData UserData
	userData.WebsocketID = websocketID
	userData.Authority = emailAccount.User.Authority
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json marshal fail",
			RequestPath: "email.login",
			Method:      "get",
		}
	}

	err = rdb.Set(ctx, strconv.Itoa(emailAccount.User.ID), jsonData, time.Minute*30).Err()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database set fail",
			RequestPath: "email.login",
			Method:      "get",
		}
	}

	return emailAccount.User.ID, HTTPStatus{
		Message:     "登录成功！",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.login",
		Method:      "get",
	}
}

/*
	有bug
	退出后信息无法保存，
*/
func (account Email) Logout(websocketID string) HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.logout",
			Method:      "delete",
		}
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.logout",
			Method:      "delete",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//未查询到账号
	if errors.Is(tx.Where("email = ?", account.Email).Find(&account).Error, gorm.ErrRecordNotFound) {
		return HTTPStatus{
			Message:     "？",
			IsError:     true,
			ErrorCode:   404,
			SubMessage:  "Hacker attack？",
			RequestPath: "email.logout",
			Method:      "delete",
		}
	}

	ctx = context.Background()
	rdb.Del(ctx, strconv.Itoa(account.UserID))

	return HTTPStatus{
		Message:     "退出成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.logout",
		Method:      "delete",
	}
}

/**/
func (account Email) Regist(websocketID string, verifiCode string) (int, HTTPStatus) {
	//检查数据库是否能够正常连接
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}

	ctx := context.Background()
	//尝试取数据
	userDataJSON, err := rdb.Get(ctx, account.Email).Result()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "验证码已过期，请重新获取验证码。",
			IsError:     true,
			ErrorCode:   412,
			SubMessage:  "verify code = null",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}

	//取数据转换json
	var userData UserData
	err = json.Unmarshal([]byte(userDataJSON), &userData)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错了。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json to UserData fail",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}

	//将取出来的验证码比较
	if userData.VerifyCode != verifiCode {
		return -1, HTTPStatus{
			Message:     "验证码出错了。",
			IsError:     true,
			ErrorCode:   412,
			SubMessage:  "user verify code error",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}

	//加入数据到mysql
	account.User.Name = verification.RandVerificationCode()
	err = mdb.Create(&account).Error
	if err != nil {
		return -1, HTTPStatus{
			Message:     "用户已有，请不要重复添加。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "email repeat",
			RequestPath: "email.regist",
			Method:      "post",
		}
	}

	//redis抹掉验证码
	rdb.Del(ctx, strconv.Itoa(account.User.ID))

	//自动登录
	rdb.Set(ctx, strconv.Itoa(account.User.ID), &UserData{
		WebsocketID: websocketID,
		Authority:   account.User.Authority,
	}, time.Minute*30)

	return account.User.ID, HTTPStatus{
		Message:     "注册成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.resign",
		Method:      "post",
	}
}

/**/
func (account Email) AuthLogin(websocketID string) HTTPStatus {
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.authlogin",
			Method:      "get",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.authlogin",
			Method:      "get",
		}
	}
	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//未查询到账号，前端需要检查并导向首页
	if errors.Is(tx.Where("email = ? AND password = ?", account.Email, account.User.Password).Find(&account).Error, gorm.ErrRecordNotFound) {
		return HTTPStatus{
			Message:     "",
			IsError:     false,
			ErrorCode:   304,
			SubMessage:  "",
			RequestPath: "email.authlogin",
			Method:      "get",
		}
	}

	//找rdb授权
	ctx = context.Background()
	res, err := rdb.Get(ctx, strconv.Itoa(account.UserID)).Result()
	if err != nil {
		return HTTPStatus{
			Message:     "亲长时间没动了。",
			IsError:     true,
			ErrorCode:   401,
			SubMessage:  "redis get error, expired",
			RequestPath: "email.authlogin",
			Method:      "get",
		}
	}
	var userData UserData

	json.Unmarshal([]byte(res), &userData)

	return HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.authlogin",
		Method:      "get",
	}
}

/**/
func (account Email) SendVerifyCode() HTTPStatus {
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}

	verifyCode := verification.RandVerificationCode()
	//验证邮箱域名是否能ping通，不能ping通则返回错误
	//TODO

	//检查用户是否存在
	var count int64
	mdb.Model(&account).Where("email = ?", account.Email).Count(&count)
	if count != 0 {
		return HTTPStatus{
			Message:     "账号已经被注册，如果忘记密码，请寻回密码",
			IsError:     true,
			ErrorCode:   0,
			SubMessage:  "账号被注册了",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}

	//验证邮箱是否发送正确
	err = email.SendMailByQQ([]string{account.Email}, "OnlineJudge", "验证码", verifyCode+"\n验证码有效时间为10分钟")
	if err != nil {
		return HTTPStatus{
			Message:     "邮箱发送错误",
			IsError:     true,
			ErrorCode:   0,
			SubMessage:  "邮箱发送错误",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}

	//将验证码等数据转换成json
	var userData UserData
	userData.VerifyCode = verifyCode
	userDataJSON, err := json.Marshal(&userData)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器发生错误",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json error",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}

	ctx := context.Background()
	//验证redis数据库是否加入成功
	err = rdb.Set(ctx, account.Email, userDataJSON, time.Minute*10).Err()
	if err != nil {
		// return err
		return HTTPStatus{
			Message:     "服务器发生错误",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis set error",
			RequestPath: "email.SendVerifyCode",
			Method:      "get",
		}
	}

	return HTTPStatus{
		Message:     "发送验证码成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.SendVerifyCode",
		Method:      "get",
	}
}

/*
	<==========================user账号相关=============================>
*/

/*
@Title
User.Login

@description
User登录模块，通常用于自动登录验证，返回用户权限给前端

@param
ID, password (int, string)

@return
authority, error (string, error)
*/
func (account User) Login(websocketID string) (int, HTTPStatus) {
	//检查数据库是否能够正常连接
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "user.login",
			Method:      "post",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "user.login",
			Method:      "post",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//TODO 加密，和数据库比较

	//未查询到账号
	var authUser User
	authUser.ID = account.ID
	if errors.Is(tx.Find(&authUser).Error, gorm.ErrRecordNotFound) {
		return -1, HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "user.login",
			Method:      "get",
		}
	}

	if account.Password != authUser.Password {
		return -1, HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "user.login",
			Method:      "get",
		}
	}

	//TODO 记录进redis中，证明已经登录过
	var userData UserData
	userData.WebsocketID = websocketID
	userData.Authority = authUser.Authority

	ctx = context.Background()
	if rdb.Set(ctx, strconv.Itoa(account.ID), userData, time.Minute*30).Err() != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦！",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis set error",
			RequestPath: "user.login",
			Method:      "get",
		}
	}

	return account.ID, HTTPStatus{
		Message:     "登录成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.login",
		Method:      "get",
	}
}

/**/
func (account User) Logout(websocketID string) HTTPStatus {
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "user.logout",
			Method:      "delete",
		}
	}

	ctx := context.Background()
	if rdb.Del(ctx, strconv.Itoa(account.ID)).Err() != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database delete fail",
			RequestPath: "user.logout",
			Method:      "delete",
		}
	}

	return HTTPStatus{
		Message:     "退出成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.logout",
		Method:      "delete",
	}
}

/*
	return: id error
*/
func (account User) Regist(websocketID string) (int, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "user.regist",
			Method:      "post",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	tx.Create(&account)

	return account.ID, HTTPStatus{
		Message:     "注册成功 ",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "",
	}
}

/*
	正在开发中
*/
func (account User) AuthLogin(websocketID string) HTTPStatus {
	// rdb, err := database.ConnectRedisDatabase(0)
	// if err != nil {
	// 	return User{}, errors.New("redis数据库连接失败！")
	// }

	// ctx := context.Background()

	// 查询是否有login
	return HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.authlogin",
		Method:      "get",
	}
}
