// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// JobQuestion is the golang structure for table job_question.
type JobQuestion struct {
	Id       string      `json:"id"        orm:"id"        ` //
	Question string      `json:"question"  orm:"question"  ` //
	JobId    string      `json:"job_id"    orm:"job_id"    ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time
}
