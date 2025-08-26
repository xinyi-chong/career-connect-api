// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Schedule is the golang structure of table schedule for DAO operations like Where/Data.
type Schedule struct {
	g.Meta        `orm:"table:schedule, do:true"`
	Id            interface{} //
	ApplicationId interface{} //
	StartTime     *gtime.Time //
	EndTime       *gtime.Time //
	Title         interface{} //
	CompanyId     interface{} //
	UserId        interface{} //
	Location      interface{} //
	Link          interface{} //
	Status        interface{} //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
