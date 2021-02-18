package web

type Database interface {
	Insert() error
	Update() error
	Query() error
	Delete() error
}

type Account interface {
	Login()
	Logout()
	Regist()
}
