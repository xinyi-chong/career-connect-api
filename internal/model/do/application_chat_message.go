// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplicationChatMessage is the golang structure of table application_chat_message for DAO operations like Where/Data.
type ApplicationChatMessage struct {
	g.Meta        `orm:"table:application_chat_message, do:true"`
	Id            interface{} //
	Name          interface{} //
	Message       interface{} //
	SenderId      interface{} //
	ApplicationId interface{} //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
