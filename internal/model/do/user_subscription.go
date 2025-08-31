// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSubscription is the golang structure of table user_subscription for DAO operations like Where/Data.
type UserSubscription struct {
	g.Meta     `orm:"table:user_subscription, do:true"`
	Id         interface{} //
	UserId     interface{} //
	UserPlanId interface{} //
	Status     interface{} //
	Expiry     *gtime.Time //
	CreateAt   *gtime.Time // Created Time
	UpdateAt   *gtime.Time // Updated Time
}
