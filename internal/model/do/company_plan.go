// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyPlan is the golang structure of table company_plan for DAO operations like Where/Data.
type CompanyPlan struct {
	g.Meta   `orm:"table:company_plan, do:true"`
	Id       interface{} //
	Name     interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
