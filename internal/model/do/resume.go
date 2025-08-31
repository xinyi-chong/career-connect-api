// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Resume is the golang structure of table resume for DAO operations like Where/Data.
type Resume struct {
	g.Meta   `orm:"table:resume, do:true"`
	Id       interface{} //
	UserId   interface{} //
	MediaId  interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
