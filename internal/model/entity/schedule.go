// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Schedule is the golang structure for table schedule.
type Schedule struct {
	Id            string      `json:"id"             orm:"id"             ` //
	ApplicationId string      `json:"application_id" orm:"application_id" ` //
	StartTime     *gtime.Time `json:"start_time"     orm:"start_time"     ` //
	EndTime       *gtime.Time `json:"end_time"       orm:"end_time"       ` //
	Title         string      `json:"title"          orm:"title"          ` //
	CompanyId     string      `json:"company_id"     orm:"company_id"     ` //
	UserId        string      `json:"user_id"        orm:"user_id"        ` //
	Location      string      `json:"location"       orm:"location"       ` //
	Link          string      `json:"link"           orm:"link"           ` //
	Status        string      `json:"status"         orm:"status"         ` //
	CreateAt      *gtime.Time `json:"create_at"      orm:"create_at"      ` // Created Time
	UpdateAt      *gtime.Time `json:"update_at"      orm:"update_at"      ` // Updated Time

	User    *User    `json:"user"           orm:"user, with:id=user_id"`       //
	Company *Company `json:"company"        orm:"company, with:id=company_id"` //
}
