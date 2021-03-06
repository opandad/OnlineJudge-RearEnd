package web

import "time"

/*
	<================mysql model====================>
*/
type ContestsHasProblem struct {
	ContestsId int `json:"contestsID"`
	ProblemsId int `json:"problemsID"`
}

type Contest struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	ContestInfo string    `json:"contestInfo` //json
}

type Email struct {
	Email  string `json:"email"`
	UserID int    `json:"userID"`
	User   User   `json:"user"`
}

type Language struct {
	ID       int    `json:"id"`
	Language string `json:"language"`
	RunCmd   string `json:"runCmd"`
}

type Problem struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"` //json
	IsHideToUser   bool   `json:"isHideToUser"`
	IsRobotProblem bool   `json:"isRobotProblem"`
	JudgeerInfo    string `json:"judggerInfo"` //json
}

type Submit struct {
	ID          int       `json:"id"`
	SubmitState string    `json:"submitState"`
	RunTime     int       `json:"runTime"`
	SubmitTime  time.Time `json:"submitTime"`
	ProblemId   int       `json:"problemsID"`
	ContestId   int       `json:"contestID"`
	UserId      int       `json:"userID"`
	LanguageId  int       `json:"languageID"`
	IsError     bool      `json:"isError"`
	SubmitInfo  string    `json:"submitInfo"`
}

type UsersJoinContest struct {
	UsersId    int `json:"userID"`
	ContestsId int `json:"contestID"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Authority string `json:"authority"`
	UserInfo  string `json:"userInfo"` //json
}

type ContestsSupportLanguage struct {
	ContestsId  int `json:"contestsID"`
	LanguagesId int `json:"languagesID"`
}

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
}

/*
	<========================front end model===============================>
*/

type FrontEndData struct {
	WebsocketID string     `json:"websocketID"`
	HTTPStatus  HTTPStatus `json:"httpStatus"`
	Data        struct {
		/*
			用户相关
		*/
		Email      []Email `json:"email"`
		Team       []Team  `json:"team"`
		User       []User  `json:"user"`
		VerifyCode string  `json:"verifyCode"`

		/*
			其他东西
		*/
		Problem             []Problem            `json:"problem"`
		Contest             []Contest            `json:"contest"`
		Language            []Language           `json:"language"`
		Submit              []Submit             `json:"submit"`
		ContestsHasProblems []ContestsHasProblem `json:"contestsHasProblems"`
		UsersJoinContests   []UsersJoinContest   `json:"usersJoinContests"`
		Page                struct {
			PageSize  int `json:"pageSize"`
			PageIndex int `json:"pageIndex"`
		} `json:"page"`
	} `json:"data"`
}

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

// <====================== end ======================>
