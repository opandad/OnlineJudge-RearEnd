package web

/*
	<=============================mysql model===================================>
	如有问题看mysql数据库设计
	ContestsHasProblem
	ContestsSupportLanguage
	UsersJoinContest
*/

type ContestsHasProblem struct {
	ContestId int `json:"contestsID"`
	ProblemId int `json:"problemsID"`
}

type ContestInfo struct {
}

type Contest struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	StartTime   string      `json:"startTime"`
	EndTime     string      `json:"endTime"`
	ContestInfo ContestInfo `json:"contestInfo` //json
	Users       []User      `gorm:"many2many:users_join_contests;"`
	Languages   []Language  `gorm:"many2many:contests_support_languages;"`
	Problems    []Problem   `gorm:"many2many:contests_has_problems;"`
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
	ProblemPath        string             `json:"problemPath"`
	ProblemJudgeConfig ProblemJudgeConfig `json:"problemConfig"`
}

type Problem struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description ProblemDescription `json:"description"`

	IsHideToUser   bool               `json:"isHideToUser"`
	IsRobotProblem bool               `json:"isRobotProblem"`
	JudgeerInfo    ProblemJudgeerInfo `json:"judgeerInfo"`
	Contests       []Contest          `gorm:"many2many:contests_has_problems;"`
}

type ProblemJudgeConfig struct {
	TestCase      []TestCase `json:"test_cases"`
	TimeLimit     int        `json:"time_limit"`
	MemoryLimit   int        `json:"memory_limit"`
	RealTimeLimit int        `json:"real_time_limit"`
	FileSizeLimit int        `json:"file_size_limit"`
	UID           int        `json:"uid"`
	StrictMode    bool       `json:"strict_mode"`
	// SpecialJudge  SpecialJudge `json:"special_judge"`
}

type SubmitInfo struct {
	CodeFileName string `json:"codeFileName"`
}

type Submit struct {
	ID          int        `json:"id"`
	SubmitState string     `json:"submitState"`
	SubmitTime  string     `json:"submitTime"`
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

type UsersJoinContest struct {
	UserId    int `json:"userID"`
	ContestId int `json:"contestID"`
}

type UserInfo struct {
	Phone string `json:"phone"`
	QQ    string `json:"qq"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Authority string    `json:"authority"`
	UserInfo  UserInfo  `json:"userInfo"`
	Contests  []Contest `gorm:"many2many:users_join_contests;"`
}

type Team struct {
	Team   string `json:"team"`
	UserID int    `json:"userID"`
	User   User   `json:"user"`
}

type Email struct {
	Email  string `json:"email"`
	UserID int    `json:"userID"`
	User   User   `json:"user"`
}

type ContestsSupportLanguage struct {
	ContestId  int `json:"contestsID" gorm:"primaryKey"`
	LanguageId int `json:"languagesID" gorm:"primaryKey"`
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
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	Authority   string `json:"authority"`
	SnowflakeID string `json:"snowflakeID"`
	VerifyCode  string `json:"verifyCode"`
}

//<========================= judger =================>
type TestCase struct {
	Handle  string `json:"handle"`
	Name    string `json:"name"`
	Input   string `json:"input"`
	Output  string `json:"output"`
	Enabled bool   `json:"enabled"`
}

// type SpecialJudge struct {
// 	Mode               int    `json:"mode"`
// 	Checker            string `json:"checker"`
// 	RedirectProgramOut bool   `json:"redirect_program_out"`
// 	TimeLimit          int    `json:"time_limit"`
// 	MemoryLimit        int    `json:"memory_limit"`
// 	UseTestLib         bool   `json:"use_testlib"`
// 	CheckerCases       string `json:"checker_cases"`
// }

// <====================== end ======================>
