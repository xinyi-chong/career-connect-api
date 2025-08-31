// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPlan is the golang structure of table user_plan for DAO operations like Where/Data.
type UserPlan struct {
	g.Meta   `orm:"table:user_plan, do:true"`
	Id       interface{} //
	Name     interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
