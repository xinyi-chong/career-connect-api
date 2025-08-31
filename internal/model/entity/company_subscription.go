// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanySubscription is the golang structure for table company_subscription.
type CompanySubscription struct {
	Id            string      `json:"id"              orm:"id"              ` //
	CompanyId     string      `json:"company_id"      orm:"company_id"      ` //
	CompanyPlanId string      `json:"company_plan_id" orm:"company_plan_id" ` //
	Status        string      `json:"status"          orm:"status"          ` //
	Expiry        *gtime.Time `json:"expiry"          orm:"expiry"          ` //
	CreateAt      *gtime.Time `json:"create_at"       orm:"create_at"       ` // Created Time
	UpdateAt      *gtime.Time `json:"update_at"       orm:"update_at"       ` // Updated Time

	CompanyPlan *CompanyPlan `json:"company_plan"   orm:"with:id=company_plan_id"` //
}
