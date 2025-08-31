// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure for table media.
type Media struct {
	Id       string      `json:"id"        orm:"id"        ` //
	Url      string      `json:"url"       orm:"url"       ` //
	Key      string      `json:"key"       orm:"key"       ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time
}
