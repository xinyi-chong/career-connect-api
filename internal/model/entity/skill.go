// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Skill is the golang structure for table skill.
type Skill struct {
	Id       string      `json:"id"        orm:"id"        ` //
	UserId   string      `json:"user_id"   orm:"user_id"   ` //
	Name     string      `json:"name"      orm:"name"      ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time
}
