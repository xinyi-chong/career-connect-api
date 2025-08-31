// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Job is the golang structure for table job.
type Job struct {
	Id            string     `json:"id"              orm:"id"              ` //
	CreatedBy     string     `json:"created_by"      orm:"created_by"      ` //
	CreatedByType string     `json:"created_by_type" orm:"created_by_type" ` //
	UpdatedBy     string     `json:"updated_by"      orm:"updated_by"      ` //
	UpdatedByType string     `json:"updated_by_type" orm:"updated_by_type" ` //
	Title         string     `json:"title"           orm:"title"           ` //
	CompanyId     string     `json:"company_id"      orm:"company_id"      ` //
	Tag           string     `json:"tag"             orm:"tag"             ` //
	Description   string     `json:"description"     orm:"description"     ` //
	Level         string     `json:"level"           orm:"level"           ` //
	Salary        string     `json:"salary"          orm:"salary"          ` //
	PostedAt      gtime.Time `json:"posted_at"       orm:"posted_at"       ` //
	Location      string     `json:"location"        orm:"location"        ` //
	IsRemote      bool       `json:"is_remote"       orm:"is_remote"       ` //
	IsHybrid      bool       `json:"is_hybrid"       orm:"is_hybrid"       ` //
	Expiry        gtime.Time `json:"expiry"          orm:"expiry"          ` //
	// JobQuestionId string      `json:"job_question_id" orm:"job_question_id" ` //
	Status   string      `json:"status"          orm:"status"          ` //
	CreateAt *gtime.Time `json:"create_at"       orm:"create_at"       ` // Created Time
	UpdateAt *gtime.Time `json:"update_at"       orm:"update_at"       ` // Updated Time

	// Company  *Company    `json:"company" orm:"company, with:id=company_id"` //
	Applications []*Application `json:"applications" orm:"applications, with:job_id=id"` //
	JobQuestion  *JobQuestion   `json:"job_question" orm:"job_question, with:job_id=id"` //
}
