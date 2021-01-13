package models

type User struct {
	ID        int
	Email     string
	Name      string
	Password  string
	Authority string
	UserInfo  string
}

func LoginUseEmail() {
	// fmt.Println("login use email")
	// db := database.GetDatabaseConnection()

	// sqlDB, err := db.DB()

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for true {
	// 	fmt.Println(sqlDB.Ping())
	// }

}

func Register() {

}

func ForgetPassword() {

}

func LoginUseWechat() {

}
