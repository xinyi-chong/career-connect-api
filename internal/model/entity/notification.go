// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure for table notification.
type Notification struct {
	Id          string      `json:"id"           orm:"id"           ` //
	RecipientId string      `json:"recipient_id" orm:"recipient_id" ` //
	Redirect    string      `json:"redirect"     orm:"redirect"     ` //
	Title       string      `json:"title"        orm:"title"        ` //
	Description string      `json:"description"  orm:"description"  ` //
	Status      string      `json:"status"       orm:"status"       ` //
	CreateAt    *gtime.Time `json:"create_at"    orm:"create_at"    ` // Created Time
	UpdateAt    *gtime.Time `json:"update_at"    orm:"update_at"    ` // Updated Time
}
