// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SuggestSkill is the golang structure of table suggest_skill for DAO operations like Where/Data.
type SuggestSkill struct {
	g.Meta   `orm:"table:suggest_skill, do:true"`
	Id       interface{} //
	Name     interface{} //
	Category interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
