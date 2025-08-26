// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notification is the golang structure of table notification for DAO operations like Where/Data.
type Notification struct {
	g.Meta      `orm:"table:notification, do:true"`
	Id          interface{} //
	RecipientId interface{} //
	Redirect    interface{} //
	Title       interface{} //
	Description interface{} //
	Status      interface{} //
	CreateAt    *gtime.Time // Created Time
	UpdateAt    *gtime.Time // Updated Time
}
