// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplicationChatMessage is the golang structure for table application_chat_message.
type ApplicationChatMessage struct {
	Id            string      `json:"id"             orm:"id"             ` //
	Name          string      `json:"name"           orm:"name"           ` //
	Message       string      `json:"message"        orm:"message"        ` //
	SenderId      string      `json:"sender_id"      orm:"sender_id"      ` //
	ApplicationId string      `json:"application_id" orm:"application_id" ` //
	CreateAt      *gtime.Time `json:"create_at"      orm:"create_at"      ` // Created Time
	UpdateAt      *gtime.Time `json:"update_at"      orm:"update_at"      ` // Updated Time
}
