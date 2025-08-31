// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Account is the golang structure for table account.
type Account struct {
	Id       string      `json:"id"        orm:"id"        ` // Account ID
	Email    string      `json:"email"     orm:"email"     ` // Email
	Password string      `json:"password"  orm:"password"  ` // Password
	Status   string      `json:"status"    orm:"status"    ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time

	User    *User    `json:"user"      orm:"user, with:account_id=id"`    //
	Company *Company `json:"company"   orm:"company, with:account_id=id"` //
}
