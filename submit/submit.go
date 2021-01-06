package submit

import "time"

type Submit struct {
	id            int
	submit_state  string
	language      string
	run_time      time.Time
	submit_time   time.Time
	problems_id   int
	contest_id    int
	language_name string
	user_id       int
}

func QuerySubmit() {
	// var submit Submit

	// db := database.GetDatabaseConnection()

	// db.Where("id = ?", 1).First(&submit)

	// fmt.Println(submit.id, submit.submit_state)
}
