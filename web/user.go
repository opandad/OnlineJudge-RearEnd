package web

/*
	<==========================email账号相关==========================>
*/

/*
正在开发

@Title
User.Login

@description
User登录模块，通常用于自动登录验证，返回用户权限给前端

@param
emailAccount, password (string, string)

@return
userID, authority, error (int, string, error)
*/
func (account Email) Login() (int, string, error) {

	return -1, "user", nil
}

/**/
func (account Email) Logout() error {
	return nil
}

/**/
func (account Email) Regist() error {
	return nil
}

/**/
func (account Email) AuthLogin() error {
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
func (account User) Login() (string, error) {
	return "user", nil
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
