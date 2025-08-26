// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Certificate is the golang structure for table certificate.
type Certificate struct {
	Id       string      `json:"id"        orm:"id"        ` // ID
	UserId   string      `json:"user_id"   orm:"user_id"   ` // User ID
	Name     string      `json:"name"      orm:"name"      ` // User email
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time
}
