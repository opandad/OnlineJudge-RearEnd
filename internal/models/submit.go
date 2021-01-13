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

func QuerySubmit() {
	// var submit Submit

	// db := database.GetDatabaseConnection()

	// db.Where("id = ?", 1).First(&submit)

	// fmt.Println(submit.id, submit.submit_state)
}
