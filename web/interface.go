package web

//<==========================接口代码===================================>

/*
	数据库一些相关操作
*/
type Database interface {
	Insert() error
	Update() error
	Select() error
	Delete() error
}

/*
	账号操作
	model: User, Email
*/
type Account interface {
	Login() (string, error) //返回websocketID，权限，错误
	Logout() error
	Regist() error
	AuthLogin() error
}
