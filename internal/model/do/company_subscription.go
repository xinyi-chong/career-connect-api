// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanySubscription is the golang structure of table company_subscription for DAO operations like Where/Data.
type CompanySubscription struct {
	g.Meta        `orm:"table:company_subscription, do:true"`
	Id            interface{} //
	CompanyId     interface{} //
	CompanyPlanId interface{} //
	Status        interface{} //
	Expiry        *gtime.Time //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
