package models

import "time"

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
