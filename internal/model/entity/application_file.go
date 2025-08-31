// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplicationFile is the golang structure for table application_file.
type ApplicationFile struct {
	Id            string      `json:"id"             orm:"id"             ` //
	ApplicationId string      `json:"application_id" orm:"application_id" ` //
	MediaId       string      `json:"media_id"       orm:"media_id"       ` //
	FileType      string      `json:"file_type"      orm:"file_type"      ` //
	CreateAt      *gtime.Time `json:"create_at"      orm:"create_at"      ` // Created Time
	UpdateAt      *gtime.Time `json:"update_at"      orm:"update_at"      ` // Updated Time

	Files       []*Media    `json:"files"          orm:"files, with:id=media_id"`             //
	Application Application `json:"application"    orm:"application, with:id=application_id"` //
}
