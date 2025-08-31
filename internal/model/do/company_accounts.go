// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyAccounts is the golang structure of table company_accounts for DAO operations like Where/Data.
type CompanyAccounts struct {
	g.Meta    `orm:"table:company_accounts, do:true"`
	UserId    interface{} //
	CompanyId interface{} //
	RoleId    interface{} //
	CreateAt  *gtime.Time // Created Time
	UpdateAt  *gtime.Time // Updated Time
}
