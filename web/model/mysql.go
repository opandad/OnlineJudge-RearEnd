package model

import "time"

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
	Email  string `gorm:"primaryKey"`
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
	ID        int
	Name      string
	Password  string
	Authority string
	UserInfo  string
}
