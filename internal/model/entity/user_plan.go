// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPlan is the golang structure for table user_plan.
type UserPlan struct {
	Id       string      `json:"id"        orm:"id"        ` //
	Name     string      `json:"name"      orm:"name"      ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time
}
