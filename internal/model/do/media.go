// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Media is the golang structure of table media for DAO operations like Where/Data.
type Media struct {
	g.Meta   `orm:"table:media, do:true"`
	Id       interface{} //
	Url      interface{} //
	Key      interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
