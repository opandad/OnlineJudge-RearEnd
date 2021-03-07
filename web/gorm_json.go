package web

import (
	"database/sql/driver"
	"encoding/json"
)

//负责实现gorm读取json所需要的alue和scan

/*
	开发列表
	ProblemDescription
	ProblemJudgeerInfo
	SubmitInfo
	UserInfo
*/

/*
<==============ProblemDescription==============>
*/
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (c *ProblemDescription) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (c ProblemDescription) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
<==============ProblemJudgeerInfo==============>
*/
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (c *ProblemJudgeerInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (c ProblemJudgeerInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
<==============SubmitInfo==============>
*/
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (c *SubmitInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (c SubmitInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

/*
<==============UserInfo==============>
*/
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (c *UserInfo) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (c UserInfo) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
