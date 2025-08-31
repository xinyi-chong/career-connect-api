// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id        string      `json:"id"         orm:"id"         ` //
	RoleId    string      `json:"role_id"    orm:"role_id"    ` //
	FeatureId string      `json:"feature_id" orm:"feature_id" ` //
	Allow     int         `json:"allow"      orm:"allow"      ` //
	CreateAt  *gtime.Time `json:"create_at"  orm:"create_at"  ` // Created Time
	UpdateAt  *gtime.Time `json:"update_at"  orm:"update_at"  ` // Updated Time

	// Role 			*Role       `json:"role" orm:"role, with:id=role_id"` //
	Feature *Feature `json:"feature"    orm:"feature, with:id=feature_id"` //
}
