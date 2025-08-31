// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Experience is the golang structure for table experience.
type Experience struct {
	Id            string      `json:"id"             orm:"id"             ` //
	UserId        string      `json:"user_id"        orm:"user_id"        ` //
	StartDate     gtime.Time  `json:"start_date"     orm:"start_date"     ` //
	EndDate       *gtime.Time `json:"end_date"       orm:"end_date"       ` //
	IsPresent     bool        `json:"is_present"     orm:"is_present"     ` //
	Description   *string     `json:"description"    orm:"description"    ` //
	Title         string      `json:"title"          orm:"title"          ` //
	CompanyId     *string     `json:"company_id"     orm:"company_id"     ` //
	CompanyString string      `json:"company_string" orm:"company_string" ` //
	CreateAt      *gtime.Time `json:"create_at"      orm:"create_at"      ` // Created Time
	UpdateAt      *gtime.Time `json:"update_at"      orm:"update_at"      ` // Updated Time

	Company *Company `json:"company"        orm:"company, with:id=company_id"` //
}
