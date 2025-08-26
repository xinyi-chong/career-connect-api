// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Application is the golang structure of table application for DAO operations like Where/Data.
type Application struct {
	g.Meta     `orm:"table:application, do:true"`
	Id         interface{} //
	JobId      interface{} //
	UserId     interface{} //
	Answer     interface{} //
	ApplyAt    *gtime.Time //
	ActivityId interface{} //
	CreateAt   *gtime.Time // Created Time
	UpdateAt   *gtime.Time // Updated Time
}
