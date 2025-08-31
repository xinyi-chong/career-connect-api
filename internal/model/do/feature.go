// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Feature is the golang structure of table feature for DAO operations like Where/Data.
type Feature struct {
	g.Meta   `orm:"table:feature, do:true"`
	Id       interface{} //
	Name     interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
