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
	ContestInfo string    `json:"contestInfo`
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
	Description    string `json:"description"`
	IsHideToUser   bool   `json:"isHideToUser"`
	IsRobotProblem bool   `json:"isRobotProblem"`
	JudgeerInfo    string `json:"judggerInfo"`
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
	UserInfo  string `json:"userInfo"`
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
type UserOnlineData struct {
	WebsocketID string `json:"websocketID"`
	VerifyCode  string `json:"verifyCode"`
}

/*
	<========================front end model===============================>
*/

type FrontEndData struct {
	WebsocketID string `json:"websocketID"`
	Message     string `json:"msg"`
	IsError     bool   `json:"isError"`
	ErrorCode   int    `json:"errorCode"`
	SubMessage  string `json:"subMsg"`
	RequestPath string `json:"requestPath"`
	Method      string `json:"method"`
	Data        struct {
		/*
			负责登录、验证时使用
		*/
		Account struct {
			ID         int    `json:"id"`
			Account    string `json:"account"`
			Password   string `json:"password"`
			Authority  string `json:"authority"`
			VerifyCode string `json:"verifyCode"`
		} `json:"account"`

		/*
			负责修改、删除用户时使用
		*/
		Email []Email `json:"email"`
		User  []User  `json:"user"`

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
