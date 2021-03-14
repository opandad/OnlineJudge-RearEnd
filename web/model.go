package web

import (
	"time"
)

/*
	<=============================mysql model===================================>
	如有问题看mysql数据库设计
	ContestsHasProblem
	ContestsSupportLanguage
	UsersJoinContest
*/

// type ContestsHasProblem struct {
// 	ContestsId int     `json:"contestsID"`
// 	Contest    Contest `json:"contest"`
// 	ProblemsId int     `json:"problemsID"`
// 	Problem    Problem `json:"problem"`
// }

type ContestInfo struct {
}

type Contest struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	StartTime   time.Time   `json:"startTime"`
	EndTime     time.Time   `json:"endTime"`
	ContestInfo ContestInfo `json:"contestInfo` //json
	Users       []User      `gorm:"many2many:users_join_contests;"`
	Languages   []Language  `gorm:"many2many:contests_support_languages;"`
	Problems    []Problem   `gorm:"many2many:contests_has_problems;"`
}

type Email struct {
	Email  string `json:"email"`
	UserID int    `json:"userID"`
	User   User   `json:"user"`
}

type Language struct {
	ID       int       `json:"id"`
	Language string    `json:"language"`
	Contests []Contest `gorm:"many2many:contests_support_languages;"`
}

type ProblemDescription struct {
	ProblemDescription string `json:"problemDescription"`
	InputDescription   string `json:"inputDescription"`
	OutputDescription  string `json:"outputDescription"`
	InputCase          string `json:"inputCase"`
	OutputCase         string `json:"outputCase"`
	Tips               string `json:"tips"`
	TimeLimit          int    `json:"timeLimit"`
	MemoryLimit        int    `json:"memoryLimit"`
	RealTimeLimit      int    `json:"realTimeLimit"`
	FileSizeLimit      int    `json:"fileSizeLimit"`
}

type ProblemJudgeerInfo struct {
	ProblemPath string `json:"problemPath"`
}

type Problem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// Description string `json:"Description"` //json
	Description ProblemDescription `json:"description"`

	IsHideToUser   bool `json:"isHideToUser"`
	IsRobotProblem bool `json:"isRobotProblem"`
	// JudgeerInfo    string `json:"judggerInfo"` //json
	JudgeerInfo ProblemJudgeerInfo `json:"judgeerInfo"`
	Contests    []Contest          `gorm:"many2many:contests_has_problems;"`
}

type SubmitInfo struct {
	CodeFileName string `json:"codeFileName"`
}

type Submit struct {
	ID          int        `json:"id"`
	SubmitState string     `json:"submitState"`
	SubmitTime  time.Time  `json:"submitTime"`
	ProblemId   int        `json:"problemID"`
	Problem     Problem    `json:"problem"`
	ContestId   int        `json:"contestID"`
	UserId      int        `json:"userID"`
	User        User       `json:"user"`
	LanguageId  int        `json:"languageID"`
	IsError     bool       `json:"isError"`
	SubmitCode  string     `json:"submitCode"`
	SubmitInfo  SubmitInfo `json:"submitInfo"`
}

// type EntryInfo struct {
// }

// type UsersJoinContest struct {
// 	UsersId    int       `json:"userID"`
// 	User       User      `json:"user"`
// 	ContestsId int       `json:"contestID"`
// 	Contest    Contest   `json:"contest"`
// 	EntryInfo  EntryInfo `json:"entryInfo"`
// }

type UserInfo struct {
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Authority string    `json:"authority"`
	UserInfo  UserInfo  `json:"userInfo"`
	Contests  []Contest `gorm:"many2many:users_join_contests;"`
}

// type ContestsSupportLanguage struct {
// 	ContestsId  int      `json:"contestsID" gorm:"primaryKey"`
// 	Contest     Contest  `json:"contest"`
// 	LanguagesId int      `json:"languagesID" gorm:"primaryKey"`
// 	Language    Language `json:"language"`
// }

type Team struct {
	Team   string `json:"team"`
	UserID int    `json:"userID"`
	User   User   `json:"user"`
}

/*
	<=========================redis model=================================>
*/

/*
	db:0
	负责存储验证码信息
	key: account
	value: struct
*/
type UserData struct {
	WebsocketID string `json:"websocketID"`
	VerifyCode  string `json:"verifyCode"`
	Authority   string `json:"authority"`
	Password    string `json:"password"`
}

/*
	<========================front end model===============================>
*/

type Page struct {
	PageSize  int   `json:"pageSize"`
	PageIndex int   `json:"pageIndex"`
	Total     int   `json:"total"`
	Total64   int64 `json:"total64"`
}

// type FrontEndData struct {
// 	WebsocketID string     `json:"websocketID"`
// 	HTTPStatus  HTTPStatus `json:"httpStatus"`
// 	Data        struct {
// 		/*
// 			用户相关
// 		*/
// 		LoginInfo LoginInfo `json:"loginInfo"`
// 		Email     []Email   `json:"email"`
// 		Team      []Team    `json:"team"`
// 		User      []User    `json:"user"`

// 		/*
// 			其他东西
// 		*/
// 		Problem             []Problem            `json:"problem"`
// 		Contest             []Contest            `json:"contest"`
// 		Language            []Language           `json:"language"`
// 		Submit              []Submit             `json:"submit"`
// 		ContestsHasProblems []ContestsHasProblem `json:"contestsHasProblems"`
// 		UsersJoinContests   []UsersJoinContest   `json:"usersJoinContests"`
// 		Page                struct {
// 			PageSize  int `json:"pageSize"`
// 			PageIndex int `json:"pageIndex"`
// 		} `json:"page"`
// 	} `json:"data"`
// }

/*
	<==== 无错误填写模板 ====>
	HTTPStatus{
		Message: "",
		IsError: false,
		ErrorCode: 0,
		SubMessage: "",
		RequestPath: "",
		Method: "",
	}

	<==== 错误填写模板 ====>
	HTTPStatus{
		Message: "",
		IsError: true,
		ErrorCode: ,
		SubMessage: "",
		RequestPath: "",
		Method: "",
	}

	###########################
	类restful
	method用于判断
	post：新建
	delete：删除
	put：更新全部信息
	get：取出资源
	patch：更新部分信息
	head
	options
*/
type HTTPStatus struct {
	Message     string `json:"msg"`
	IsError     bool   `json:"isError"`
	ErrorCode   int    `json:"errorCode"`
	SubMessage  string `json:"subMsg"`
	RequestPath string `json:"requestPath"` //类路径
	Method      string `json:"method"`      //废弃
}

type LoginInfo struct {
	Account     string `json:"account"`
	UserID      int    `json:"userID"`
	Password    string `json:"password"`
	Authority   string `json:"authority"`
	SnowflakeID string `json:"snowflakeID"`
	VerifyCode  string `json:"verifyCode"`
}

//<========================= judger =================>

// <====================== end ======================>
