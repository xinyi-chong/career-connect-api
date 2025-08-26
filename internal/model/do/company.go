// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Company is the golang structure of table company for DAO operations like Where/Data.
type Company struct {
	g.Meta      `orm:"table:company, do:true"`
	Id          interface{} // company ID
	AccountId   interface{} //
	Name        interface{} //
	Description interface{} //
	Industry    interface{} //
	Tag         interface{} //
	Address     interface{} //
	Website     interface{} //
	City        interface{} //
	Size        interface{} //
	Contact     interface{} //
	LogoId      interface{} //
	CreateAt    *gtime.Time // Created Time
	UpdateAt    *gtime.Time // Updated Time
}
