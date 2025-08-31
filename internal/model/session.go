package model

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/os/gtime"
)

type AccountWithoutPassword struct {
	Id       string      `json:"id"        orm:"id"        ` // Account ID
	Email    string      `json:"email"     orm:"email"     ` // Email
	Status   string      `json:"status"    orm:"status"    ` //
	CreateAt *gtime.Time `json:"create_at" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at" ` // Updated Time

	// User    *entity.User    `orm:"with:account_id=id"    ` //
	// Company *entity.Company `orm:"with:account_id=id"    ` //
}

type Permission struct {
	CompanyID  string   `json:"company_id"`
	RoleID     string   `json:"role_id"`
	FeatureIDs []string `json:"feature_ids"`
}

type SessionData struct {
	Account     AccountWithoutPassword `json:"account"`
	User        *entity.User           `json:"user"`
	Company     *entity.Company        `json:"company"`
	Permissions []*Permission          `json:"permissions"`
}
