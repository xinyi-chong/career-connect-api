// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Resume is the golang structure for table resume.
type Resume struct {
	Id       string      `json:"id"        orm:"id"        ` //
	UserId   string      `json:"user_id"   orm:"user_id"   ` //
	MediaId  string      `json:"media_id"  orm:"media_id"  ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time

	Media *Media `json:"media"     orm:"media, with:id=media_id"` //
}
