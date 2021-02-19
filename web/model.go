package web

import "time"

/*
	<================mysql model====================>
*/
type ContestsHasProblem struct {
	ContestsId int
	ProblemsId int
}

type Contest struct {
	ID          int
	Name        string
	StartTime   time.Time
	Duration    time.Time
	ContestInfo string
}

type Email struct {
	Email  string
	UserID int
	User   User
}

type Language struct {
	ID       int
	Language string
	RunCmd   string
}

type Problem struct {
	ID             int
	Name           string
	Description    string
	Accept         int
	Fail           int
	IsRobotProblem bool
	JudgeerInfo    string
}

type Submit struct {
	ID          int
	SubmitState string
	RunTime     time.Time
	SubmitTime  time.Time
	ProblemsId  int
	ContestId   int
	UserId      int
	LanguageId  int
}

type UsersJoinContest struct {
	UsersId    int
	ContestsId int
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
	ErrorCode   int    `json:"httpStatusCode"`
	RequestPath string `json:"requestPath"`
	Function    string `json:"function"`
	Data        struct {
		Email               []Email              `json:"email"`
		User                []User               `json:"user"`
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
		VerifyCode string `json:"verifyCode"`
	} `json:"data"`
}
