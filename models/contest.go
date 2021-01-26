package models

import "time"

type Contest struct {
	ID          int
	Name        string
	StartTime   time.Time
	Duration    time.Time
	ContestInfo string
}
