// POST	/Schedule	添加职位。
// PATCH	/Schedule/{Schedule_id}	更新特定的职位。
// DELETE	/Schedule/{Schedule_id}	删除特定的职位。
// GET	/Schedule/me/Schedule	检索职位列表。
// GET	/Schedule/{Schedule_id}/Schedule	检索职位列表。
// GET	/Schedule/Schedule/{Schedule_id}	检索职位列表。
// GET	/Schedule/me/Schedule	检索认证用户的职位申请。
// GET	/Schedule/{Schedule_id}/Schedule	检索职位申请。
// POST	/Schedule/question	添加职位相关问题。
// GET	/Schedule/questions	检索职位相关问题。
// GET	/Schedule/question/{question_id}	检索特定的职位相关问题。
// PATCH	/Schedule/question/{question_id}	更新特定的职位相关问题。
// DELETE	/Schedule/question/{question_id}	删除特定的职位相关问题。

package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type PostCreateScheduleByApplicationIDReq struct {
	g.Meta   `path:"/application/:application_id" method:"post" tags:"Schedule" summary:"Create Schedule By Application ID 添加特定职位申请的面试"`
	StartTime 	gtime.Time   `json:"start_time" v:"required"`
	EndTime 		gtime.Time   `json:"end_time"   v:"required"`
	Title 			string       `json:"title"      v:"required"`
	Location 		string       `json:"location"   v:"required"`
	Link 				*string      `json:"link"       v:"required"`
}

type PostCreateScheduleByApplicationIDRes struct{
	Id 				*string				`json:"id"`
}

type GetScheduleMeCompanyReq struct {
	g.Meta   `path:"/me/company" method:"get" tags:"Schedule" summary:"Get Schedule Me Company 检索认证用户公司的面试。"`
}

type GetScheduleMeCompanyRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleMeUserReq struct {
	g.Meta   `path:"/me/user" method:"get" tags:"Schedule" summary:"Get Schedule Me User 检索认证用户的面试"`
}
type GetScheduleMeUserRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleByCompanyIDReq struct {
	g.Meta   `path:"/company/:company_id" method:"get" tags:"Schedule" summary:"Get Schedule By Company ID 检索特定公司的面试。"`
}

type GetScheduleByCompanyIDRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleByApplicationIDReq struct {
	g.Meta   `path:"/application/{application_id}" method:"get" tags:"Schedule" summary:"Get Schedule By Application ID 检索特定职位申请的面试。"`

}
type GetScheduleByApplicationIDRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleByApplicationIDCompanyReq struct {
	g.Meta   `path:"/:application_id/company" method:"get" tags:"Schedule" summary:"Get Schedule By Application ID Company 根据应用程序ID检索公司的面试。"`

}
type GetScheduleByApplicationIDCompanyRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleByApplicationIDUserReq struct {
	g.Meta   `path:"/:application_id/user" method:"get" tags:"Schedule" summary:"Get Schedule By Application ID User 根据应用程序ID检索用户的面试。"`

}
type GetScheduleByApplicationIDUserRes struct {
	Schedules []*entity.Schedule	`json:"schedules"`
}

type GetScheduleByIDCompanyReq struct {
	g.Meta   `path:"/:schedule_id/company" method:"get" tags:"Schedule" summary:"Get Schedule By ID Company 根据面试ID检索特定公司的面试。"`

}
type GetScheduleByIDCompanyRes struct {
	Schedule *entity.Schedule	`json:"schedule"`
}

type GetScheduleByIDUserReq struct {
	g.Meta   `path:"/:schedule_id/user" method:"get" tags:"Schedule" summary:"Get Schedule By ID User 根据面试ID检索特定用户的面试。"`

}
type GetScheduleByIDUserRes struct {
	Schedule *entity.Schedule	`json:"schedule"`
}

type PatchUpdateScheduleByIDReq struct {
	g.Meta   `path:"/:schedule_id" method:"patch" tags:"Schedule" summary:"Update Schedule By ID 根据面试ID更新特定的面试"`
	StartTime 	gtime.Time   `json:"start_time" v:"required"`
	EndTime 		gtime.Time   `json:"end_time"   v:"required"`
	Title 			string       `json:"title"      v:"required"`
	Location 		string       `json:"location"   v:"required"`
	Link 				*string      `json:"link"       v:"required"`
	Status 		  string       `json:"status"     v:"required"`
}
type PatchUpdateScheduleByIDRes struct {}

type DeleteScheduleByIDReq struct {
	g.Meta   `path:"/:schedule_id" method:"delete" tags:"Schedule" summary:"Delete Schedule By ID 根据面试ID删除特定的面试。"`
}

type DeleteScheduleByIDRes struct {}
