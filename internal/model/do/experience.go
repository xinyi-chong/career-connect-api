// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Experience is the golang structure of table experience for DAO operations like Where/Data.
type Experience struct {
	g.Meta        `orm:"table:experience, do:true"`
	Id            interface{} //
	UserId        interface{} //
	StartDate     *gtime.Time //
	EndDate       *gtime.Time //
	IsPresent     interface{} //
	Description   interface{} //
	Title         interface{} //
	CompanyId     interface{} //
	CompanyString interface{} //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
