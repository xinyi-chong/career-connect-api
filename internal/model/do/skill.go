// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Skill is the golang structure of table skill for DAO operations like Where/Data.
type Skill struct {
	g.Meta   `orm:"table:skill, do:true"`
	Id       interface{} //
	UserId   interface{} //
	Name     interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
