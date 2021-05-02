package web

import (
	"OnlineJudge-RearEnd/api/database"
	"OnlineJudge-RearEnd/api/email"
	"OnlineJudge-RearEnd/api/excel"
	"OnlineJudge-RearEnd/api/verification"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	| Login               |   yes   |    yes    |  yes  |

	| Logout              |   yes   |    yes    |  yes  |

	| Regist              |   yes   |    yes    |  no   |

	| AuthLogin           |   yes   |    yes    |  yes  |

	| SendVerifyCode      |   yes   |    yes    |  no   |

	Class name: user

	| func name           | develop  | unit test |

	|--------------------------------------------|

	| Login               |    yes   |    yes	 |

	| Logout              |    yes   |    yes	 |

	| AuthLogin           |    yes   |    no	 |
*/

/*
	bug list
	退出直接删除，不会验证用户
	验证用户不会验证是否异地登录
	用户名唯一bug
	无批量注册功能
*/
type Account interface {
	Login(websocketID string) (int, HTTPStatus) //返回userID和HTTPStatus
	Logout(websocketID string) HTTPStatus
	// Regist(websocketID string, verifiCode string) (int, HTTPStatus) //返回userID和HTTPStatus
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
func (account Email) Login(websocketID string) (int, string, string, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return -1, "", "", HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.login",
			Method:      "",
		}
	}
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return -1, "", "", HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.login",
			Method:      "",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	var emailAccount Email
	emailAccount.Email = account.Email
	//未查询到账号
	if errors.Is(tx.Where("email = ?", emailAccount.Email).Find(&emailAccount).Error, gorm.ErrRecordNotFound) {
		return -1, "", "", HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "email.login",
			Method:      "",
		}
	}

	tx.Model(&emailAccount).Association("User").Find(&emailAccount.User)

	//TODO 加密，和数据库比较

	if emailAccount.User.Password != account.User.Password {
		return -1, "", "", HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "email.login",
			Method:      "",
		}
	}

	//数据记录并转换json
	var userData UserData
	userData.WebsocketID = websocketID
	userData.Authority = emailAccount.User.Authority
	jsonData, err := json.Marshal(userData)
	if err != nil {
		return -1, "", "", HTTPStatus{
			Message:     "服务器出错啦。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json marshal fail",
			RequestPath: "email.login",
			Method:      "",
		}
	}

	err = rdb.Set(ctx, strconv.Itoa(emailAccount.User.ID), jsonData, time.Minute*30).Err()
	if err != nil {
		return -1, "", "", HTTPStatus{
			Message:     "服务器出错啦。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database set fail",
			RequestPath: "email.login",
			Method:      "",
		}
	}

	return emailAccount.User.ID, emailAccount.User.Authority, emailAccount.User.Name, HTTPStatus{
		Message:     "登录成功！",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.login",
		Method:      "LoginByEmail",
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
			Method:      "",
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
			Method:      "",
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
			Method:      "",
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
		Method:      "Logout",
	}
}

/**/
func (account Email) Regist(websocketID string, verifiCode string) (User, HTTPStatus) {
	//检查数据库是否能够正常连接
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "email.regist",
			Method:      "",
		}
	}
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "email.regist",
			Method:      "",
		}
	}

	ctx := context.Background()
	//尝试取数据
	userDataJSON, err := rdb.Get(ctx, account.Email).Result()
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "验证码已过期，请重新获取验证码。",
			IsError:     true,
			ErrorCode:   412,
			SubMessage:  "verify code = null",
			RequestPath: "email.regist",
			Method:      "",
		}
	}

	//取数据转换json
	var userData UserData
	err = json.Unmarshal([]byte(userDataJSON), &userData)
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "服务器出错了。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json to UserData fail",
			RequestPath: "email.regist",
			Method:      "",
		}
	}

	//将取出来的验证码比较
	if userData.VerifyCode != verifiCode {
		return User{}, HTTPStatus{
			Message:     "验证码出错了。",
			IsError:     true,
			ErrorCode:   412,
			SubMessage:  "user verify code error",
			RequestPath: "email.regist",
			Method:      "",
		}
	}

	//加入数据到mysql
	account.User.Name = verification.RandVerificationCode()
	account.User.Authority = "user"
	err = mdb.Create(&account).Error
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "用户已有，请不要重复添加。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "email repeat",
			RequestPath: "email.regist",
			Method:      "",
		}
	}

	//redis抹掉验证码
	rdb.Del(ctx, strconv.Itoa(account.User.ID))

	//自动登录
	rdb.Set(ctx, strconv.Itoa(account.User.ID), &UserData{
		WebsocketID: websocketID,
		Authority:   account.User.Authority,
	}, time.Minute*30)

	return account.User, HTTPStatus{
		Message:     "注册成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.resign",
		Method:      "LoginByEmail",
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
			RequestPath: "email.authlogin",
			Method:      "",
		}
	}
	ctx := context.Background()
	tx := mdb.WithContext(ctx)

	//未查询到账号，前端需要检查并导向首页
	if errors.Is(tx.Where("email = ?", account.Email).Find(&account).Error, gorm.ErrRecordNotFound) {
		return HTTPStatus{
			Message:     "",
			IsError:     false,
			ErrorCode:   304,
			SubMessage:  "",
			RequestPath: "email.authlogin",
			Method:      "ReturnIndex",
		}
	}
	password := account.User.Password
	tx.Model(&account).Association("User").Find(&account.User)
	if password != account.User.Password {
		return HTTPStatus{
			Message:     "账号可能被盗。",
			IsError:     true,
			ErrorCode:   401,
			SubMessage:  "password error",
			RequestPath: "email.authlogin",
			Method:      "ReturnIndex",
		}
	}

	//找rdb授权
	ctx = context.Background()
	res, err := rdb.Get(ctx, strconv.Itoa(account.UserID)).Result()
	if err != nil {
		fmt.Println(err)
		return HTTPStatus{
			Message:     "亲长时间没动了。",
			IsError:     true,
			ErrorCode:   401,
			SubMessage:  "redis get error, expired",
			RequestPath: "email.authlogin",
			Method:      "ReturnIndex",
		}
	}
	var userData UserData
	json.Unmarshal([]byte(res), &userData)

	//bug
	//异地登录不会踢

	return HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "email.authlogin",
		Method:      "",
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

func (account User) GetUserInfo() (User, HTTPStatus) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return User{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "user.login",
			Method:      "",
		}
	}

	mdb.First(&account)

	return account, HTTPStatus{
		IsError:     false,
		ErrorCode:   0,
		RequestPath: "user: get user info",
	}
}

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
			Method:      "",
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
			Method:      "",
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
			Method:      "",
		}
	}

	if account.Password != authUser.Password {
		return -1, HTTPStatus{
			Message:     "账号或密码出错啦！",
			IsError:     true,
			ErrorCode:   406,
			SubMessage:  "account or password error",
			RequestPath: "user.login",
			Method:      "",
		}
	}

	//TODO 记录进redis中，证明已经登录过
	var userData UserData
	userData.WebsocketID = websocketID
	userData.Authority = authUser.Authority
	userJsonData, err := json.Marshal(userData)
	if err != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦！",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json marshal error",
			RequestPath: "user.login",
			Method:      "",
		}
	}
	ctx = context.Background()
	if rdb.Set(ctx, strconv.Itoa(account.ID), userJsonData, time.Minute*30).Err() != nil {
		return -1, HTTPStatus{
			Message:     "服务器出错啦！",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis set error",
			RequestPath: "user.login",
			Method:      "",
		}
	}

	return account.ID, HTTPStatus{
		Message:     "登录成功！",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.login",
		Method:      "LoginByUser",
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
			Method:      "",
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
			Method:      "",
		}
	}

	return HTTPStatus{
		Message:     "退出成功",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.logout",
		Method:      "Logout",
	}
}

/**/
func (account User) AuthLogin(websocketID string) HTTPStatus {
	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "user.authlogin",
			Method:      "",
		}
	}
	ctx := context.Background()

	result, err := rdb.Get(ctx, strconv.Itoa(account.ID)).Result()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "user.authlogin",
			Method:      "",
		}
	}
	var userData UserData
	json.Unmarshal([]byte(result), &userData)
	// 查询是否有login
	if userData.WebsocketID == "" {
		return HTTPStatus{
			Message:     "登录已过期。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "登录已过期",
			RequestPath: "user.authlogin",
			Method:      "",
		}
	}

	if userData.WebsocketID != websocketID {
		return HTTPStatus{
			Message:     "已在其他地方登录。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "已在其他地方登录",
			RequestPath: "user.authlogin",
			Method:      "Logout",
		}
	}

	// 返回无错误
	return HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "user.authlogin",
		Method:      "",
	}
}

// <======================== login info =============================>

func (loginInfo LoginInfo) AuthLogin() HTTPStatus {
	if loginInfo.UserID == 0 || loginInfo.SnowflakeID == "" {
		return HTTPStatus{
			IsError: false,
		}
	}

	rdb, err := database.ConnectRedisDatabase(0)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "error",
		}
	}

	ctx := context.Background()
	result, err := rdb.Get(ctx, strconv.Itoa(loginInfo.UserID)).Result()
	if err != nil {
		return HTTPStatus{
			Message:     "登录过期",
			IsError:     false,
			ErrorCode:   500,
			SubMessage:  "登录过期",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "Logout",
		}
	}
	var authLogin LoginInfo
	err = json.Unmarshal([]byte(result), &authLogin)
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错了",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "json unmarshal error",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "error",
		}
	}

	if loginInfo.Authority != authLogin.Authority {
		return HTTPStatus{
			Message:     "登录过期",
			IsError:     false,
			ErrorCode:   500,
			SubMessage:  "authority error",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "Logout",
		}
	}
	if loginInfo.SnowflakeID != authLogin.Authority {
		return HTTPStatus{
			Message:     "在其他地方登录",
			IsError:     false,
			ErrorCode:   500,
			SubMessage:  "websocket error",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "Logout",
		}
	}
	if loginInfo.Password != authLogin.Password {
		return HTTPStatus{
			Message:     "账号疑似被盗，请修改密码",
			IsError:     false,
			ErrorCode:   500,
			SubMessage:  "password error",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "Logout",
		}
	}
	return HTTPStatus{
		IsError: false,
	}
}

func (loginInfo LoginInfo) AuthAdmin() HTTPStatus {
	if loginInfo.UserID == 0 || loginInfo.SnowflakeID == "" {
		return HTTPStatus{
			IsError: false,
		}
	}

	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "redis database connect fail",
			RequestPath: "LoginInfo.AuthLogin",
			Method:      "error",
		}
	}

	ctx := context.Background()
	tx := mdb.WithContext(ctx)
	var tmpUser User
	tmpUser.ID = loginInfo.UserID
	tmpUser.Password = loginInfo.Password
	tmpUser.Authority = loginInfo.Authority

	var user User
	var count int64
	tx.Where(&tmpUser).Find(&user).Count(&count)
	if count <= 0 {
		return HTTPStatus{
			IsError:    true,
			Message:    "你不是管理员",
			SubMessage: "该用户非管理员",
			Method:     "logininfo.AuthAdmin",
		}
	} else {
		return HTTPStatus{
			IsError:    false,
			Message:    "",
			SubMessage: "",
			Method:     "logininfo.AuthAdmin",
		}
	}
}

// <============ database ===============>
func (user User) List(pageIndex int, pageSize int) ([]User, HTTPStatus, int64) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []User{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "user.list",
			Method:      "",
		}, 0
	}

	//分页查询
	if pageIndex <= 0 || pageSize <= 0 {
		return []User{}, HTTPStatus{
			Message:     "非法输入",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "page index or page size input error, error code is error",
			RequestPath: "user.list",
			Method:      "",
		}, 0
	}

	var users []User
	var count int64
	mdb.Table("users").Count(&count).Debug().Offset((pageIndex-1)*pageSize).Limit(pageSize).Select("id", "name").Find(&users)
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

	return users, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "Get user list",
	}, count
}

func (team Team) List(pageIndex int, pageSize int) ([]Team, HTTPStatus, Page) {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return []Team{}, HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "user.list",
			Method:      "",
		}, Page{}
	}

	//分页查询
	if pageIndex <= 0 || pageSize <= 0 {
		return []Team{}, HTTPStatus{
			Message:     "非法输入",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "page index or page size input error, error code is error",
			RequestPath: "user.list",
			Method:      "",
		}, Page{}
	}

	var teams []Team
	var count int64
	var page Page
	mdb.Table("teams").Count(&count).Debug().Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&teams)
	page.PageIndex = pageIndex
	page.PageSize = pageSize
	page.Total64 = count
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

	return teams, HTTPStatus{
		Message:     "",
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "",
		Method:      "Get user list",
	}, page
}

func (team Team) AddTeamsByExcel(filePath string) HTTPStatus {
	data, err := excel.ReadTeam(filePath)
	if err != nil {
		return HTTPStatus{
			Message:     "文件读取失败",
			IsError:     true,
			RequestPath: "team.add teams by excel",
		}
	}
	// fmt.Println(data)

	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "team.add team by excel",
			Method:      "",
		}
	}

	var count int64
	mdb.Table("teams").Count(&count)
	teams := make([]Team, len(data)-1)
	for i := 1; i < len(data); i++ {
		teams[i-1].User.Name = data[i][0]
		teams[i-1].User.UserInfo.Phone = data[i][1]
		teams[i-1].User.UserInfo.QQ = data[i][2]
		teams[i-1].User.Authority = "user"
		teams[i-1].User.Password = verification.RandVerificationCode()
		teams[i-1].Team = "team" + strconv.Itoa(int(count)+i)
	}
	mdb.Create(&teams)

	fmt.Println(teams)

	return HTTPStatus{
		Message:     "成功，号码从" + strconv.Itoa(int(count)+1) + "开始，截止至" + strconv.Itoa(int(count)+len(data)-1),
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "team.add team by excel",
		Method:      "",
	}
}

func (team Team) AddTeamsByHTML(count int) HTTPStatus {
	mdb, err := database.ReconnectMysqlDatabase()
	if err != nil {
		return HTTPStatus{
			Message:     "服务器出错啦，请稍后重新尝试。",
			IsError:     true,
			ErrorCode:   500,
			SubMessage:  "mysql database connect fail",
			RequestPath: "team.add team by excel",
			Method:      "",
		}
	}

	var cnt int64
	mdb.Table("teams").Count(&cnt)
	teams := make([]Team, count)
	for i := 1; i <= count; i++ {
		teams[i-1].User.Name = "队伍" + strconv.Itoa(i+int(cnt))
		teams[i-1].User.Authority = "user"
		teams[i-1].User.Password = verification.RandVerificationCode()
		teams[i-1].Team = "team" + strconv.Itoa(int(cnt)+i)
		fmt.Print("正在添加第%d个", i)
	}
	mdb.Create(&teams)

	fmt.Println(teams)

	return HTTPStatus{
		Message:     "成功，号码从" + strconv.Itoa(int(cnt)+1) + "开始，截止至" + strconv.Itoa(int(cnt)+count),
		IsError:     false,
		ErrorCode:   0,
		SubMessage:  "",
		RequestPath: "team.add team by excel",
		Method:      "",
	}
}
