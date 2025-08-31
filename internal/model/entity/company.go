// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Company is the golang structure for table company.
type Company struct {
	Id          string      `json:"id"          orm:"id"          `    // Company ID
	AccountId   string      `json:"account_id"  orm:"account_id"  `    //
	Name        string      `json:"name"        orm:"name"        `    //
	Description string      `json:"description" orm:"description" `    //
	Industry    string      `json:"industry"    orm:"industry"    `    //
	Tag         string      `json:"tag"         orm:"tag"         `    //
	Address     string      `json:"address"     orm:"address"     `    //
	Website     string      `json:"website"     orm:"website"     `    //
	City        string      `json:"city"        orm:"city"        `    //
	Size        string      `json:"size"        orm:"size"        `    //
	Contact     string      `json:"contact"     orm:"contact"     `    //
	LogoId      string      `json:"logo_id"     orm:"logo_id"        ` //
	CreateAt    *gtime.Time `json:"create_at"   orm:"create_at"   `    // Created Time
	UpdateAt    *gtime.Time `json:"update_at"   orm:"update_at"   `    // Updated Time

	Logo                *Media               `json:"logo"                 orm:"logo, with:id=logo_id"`                    //
	CompanySubscription *CompanySubscription `json:"company_subscription" orm:"company_subscription, with:company_id=id"` //
	Companyaccounts     []*CompanyAccounts   `json:"company_accounts"     orm:"company_accounts, with:company_id=id"`     //
	Notifications       []*Notification      `json:"notifications"        orm:"notifications, with:recipient_id=id"`      //
}
