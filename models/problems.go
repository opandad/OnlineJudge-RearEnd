package models

type Problem struct {
	id               int
	name             string
	description      string
	accept           int
	fail             int
	is_robot_problem bool
	judgeer_info     string
}
