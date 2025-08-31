// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// JobQuestion is the golang structure of table job_question for DAO operations like Where/Data.
type JobQuestion struct {
	g.Meta   `orm:"table:job_question, do:true"`
	Id       interface{} //
	Question interface{} //
	JobId    interface{} //
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
}
