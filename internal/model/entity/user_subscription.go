// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSubscription is the golang structure for table user_subscription.
type UserSubscription struct {
	Id         string      `json:"id"           orm:"id"           ` //
	UserId     string      `json:"user_id"      orm:"user_id"      ` //
	UserPlanId string      `json:"user_plan_id" orm:"user_plan_id" ` //
	Status     string      `json:"status"       orm:"status"       ` //
	Expiry     *gtime.Time `json:"expiry"       orm:"expiry"       ` //
	CreateAt   *gtime.Time `json:"create_at"    orm:"create_at"    ` // Created Time
	UpdateAt   *gtime.Time `json:"update_at"    orm:"update_at"    ` // Updated Time

	UserPlan *UserPlan `json:"user_plan"    orm:"user_plan, with:id=user_plan_id"` //
}
