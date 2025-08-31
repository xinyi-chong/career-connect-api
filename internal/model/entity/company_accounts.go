// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyAccounts is the golang structure for table company_accounts.
type CompanyAccounts struct {
	UserId    string      `json:"user_id"    orm:"user_id"    ` //
	CompanyId string      `json:"company_id" orm:"company_id" ` //
	RoleId    string      `json:"role_id"    orm:"role_id"    ` //
	CreateAt  *gtime.Time `json:"create_at"  orm:"create_at"  ` // Created Time
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  ` // Updated Time

	// User 	 	 *User     	 `json:"user" orm:"user, with:id=user_id"` //
	// Company  *Company    `json:"company" orm:"company, with:id=company_id"` //
	Role *Role `json:"role"       orm:"role, with:id=role_id"` //
}
