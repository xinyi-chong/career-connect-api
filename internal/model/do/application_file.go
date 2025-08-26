// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApplicationFile is the golang structure of table application_file for DAO operations like Where/Data.
type ApplicationFile struct {
	g.Meta        `orm:"table:application_file, do:true"`
	Id            interface{} //
	ApplicationId interface{} //
	MediaId       interface{} //
	FileType      interface{} //
	CreateAt      *gtime.Time // Created Time
	UpdateAt      *gtime.Time // Updated Time
}
