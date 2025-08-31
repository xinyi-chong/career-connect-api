// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Education is the golang structure for table education.
type Education struct {
	Id              string      `json:"id"               orm:"id"               ` //
	UserId          string      `json:"user_id"          orm:"user_id"     `      //
	StartDate       *gtime.Time `json:"start_date"       orm:"start_date"       ` //
	EndDate         *gtime.Time `json:"end_date"         orm:"end_date"         ` //
	InstituteId     string      `json:"institute_id"     orm:"institute_id"     ` //
	InstituteString string      `json:"institute_string" orm:"institute_string" ` //
	Level           string      `json:"level"            orm:"level"            ` //
	Programme       string      `json:"programme"        orm:"programme"        ` //
	Description     string      `json:"description"      orm:"description"      ` //
	CreateAt        *gtime.Time `json:"create_at"        orm:"create_at"        ` // Created Time
	UpdateAt        *gtime.Time `json:"update_at"        orm:"update_at"        ` // Updated Time

	Institute *Company `json:"institute"        orm:"institute, with:id=institute_id"` //
}
