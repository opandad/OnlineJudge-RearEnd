package web

import (
	"database/sql/driver"
	"encoding/json"
)

/*
	负责实现gorm读取json所需要的value和scan

	开发列表
	ProblemDescription
	ProblemJudgeerInfo
	SubmitInfo
	UserInfo
	ContestInfo
*/
type GormJson interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)
}

/*
#######ProblemDescription#######
*/
func (c *ProblemDescription) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c ProblemDescription) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
#######ProblemJudgeerInfo#######
*/
func (c *ProblemJudgeerInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c ProblemJudgeerInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
#######SubmitInfo#######
*/
func (c *SubmitInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c SubmitInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
#######UserInfo#######
*/
func (c *UserInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c UserInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
   #######ContestInfo###########
*/
func (c *ContestInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

func (c ContestInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
