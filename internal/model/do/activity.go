// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activity is the golang structure of table activity for DAO operations like Where/Data.
type Activity struct {
	g.Meta   `orm:"table:activity, do:true"`
	Id       interface{} //
	Name     interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
