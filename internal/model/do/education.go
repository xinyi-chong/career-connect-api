// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Education is the golang structure of table education for DAO operations like Where/Data.
type Education struct {
	g.Meta          `orm:"table:education, do:true"`
	Id              interface{} //
	UserId          interface{} //
	StartDate       *gtime.Time //
	EndDate         *gtime.Time //
	InstituteId     interface{} //
	InstituteString interface{} //
	Level           interface{} //
	Programme       interface{} //
	Description     interface{} //
	CreateAt        *gtime.Time // Created Time
	UpdateAt        *gtime.Time // Updated Time
}
