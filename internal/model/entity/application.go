// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Application is the golang structure for table application.
type Application struct {
	Id     string `json:"id"          orm:"id"          ` //
	JobId  string `json:"job_id"      orm:"job_id"      ` //
	UserId string `json:"user_id"     orm:"user_id"     ` //
	Answer string `json:"answer"      orm:"answer"      ` //
	// ResumeId   *string     `json:"resume_id"   orm:"resume_id"   ` //
	ApplyAt    *gtime.Time `json:"apply_at"    orm:"apply_at"    ` //
	ActivityId string      `json:"activity_id" orm:"activity_id" ` //
	CreateAt   *gtime.Time `json:"create_at"   orm:"create_at"   ` // Created Time
	UpdateAt   *gtime.Time `json:"update_at"   orm:"update_at"   ` // Updated Time

	Job  *Job  `json:"job"         orm:"job, with:id=job_id"`   //
	User *User `json:"user"        orm:"user, with:id=user_id"` //
	// Resume   	 *Media      `json:"resume"      orm:"resume, with:id=resume_id"   ` // Resume
	Activity                *Activity                 `json:"activity"    orm:"activity, with:id=activity_id"   ` // Activity
	ApplicationChatMessages []*ApplicationChatMessage `json:"application_chat_messages" orm:"application_chat_messages, with:application_id=id"`
	Schedules               []*Schedule               `json:"schedules"   orm:"schedules, with:application_id=id"`
	ApplicationFiles        []*ApplicationFile        `json:"application_files"      orm:"application_files, with:application_id=id"   ` // Application Files
}
