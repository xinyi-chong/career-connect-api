// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure of table account for DAO operations like Where/Data.
type Account struct {
	g.Meta   `orm:"table:account, do:true"`
	Id       interface{} // account ID
	Email    interface{} // Email
	Password interface{} // Password
	Status   interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
