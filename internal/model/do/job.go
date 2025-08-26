// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Job is the golang structure of table job for DAO operations like Where/Data.
type Job struct {
	g.Meta        `orm:"table:job, do:true"`
	Id            interface{} //
	CreatedBy     interface{} //
	CreatedByType interface{} //
	UpdatedBy     interface{} //
	UpdatedByType interface{} //
	Title         interface{} //
	CompanyId     interface{} //
	Tag           interface{} //
	Description   interface{} //
	Level         interface{} //
	Salary        interface{} //
	PostedAt      *gtime.Time //
	Location      interface{} //
	IsRemote      interface{} //
	IsHybrid      interface{} //
	Expiry        *gtime.Time //
	Status        interface{} //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
