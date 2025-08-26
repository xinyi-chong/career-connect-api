// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Certificate is the golang structure of table certificate for DAO operations like Where/Data.
type Certificate struct {
	g.Meta   `orm:"table:certificate, do:true"`
	Id       interface{} // ID
	UserId   interface{} // user ID
	Name     interface{} // user email
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
