package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type PostCreateJobReq struct {
	g.Meta `path:"/" tags:"Job" method:"post" summary:"Create Job"`
	Title         string      `json:"title"            v:"required"`
	CompanyID     *string     `json:"company_id"`
	Tag           string      `json:"tag"              v:"required"`
	Description	  string      `json:"description"      v:"required"`
	Level	  			string      `json:"level"            v:"required"`
	Salary	  		string      `json:"salary"           v:"required"`
	PostedAt	  	gtime.Time  `json:"posted_at"        v:"required"`
	Location	  	string      `json:"location"         v:"required"`
	IsRemote	  	bool        `json:"is_remote"        v:"required"`
	IsHybrid	  	bool        `json:"is_hybrid"        v:"required"`
	Expiry	  		gtime.Time  `json:"expiry"           v:"required"`
	JobQuestion		*string     `json:"job_question"`
}

type PostCreateJobRes struct {
	Id					*string  		`json:"id"`
}

type GetJobsAllMeReq struct {
	g.Meta `path:"/all/me" tags:"Job" method:"get" summary:"Get Jobs All Me (User)"`
}

type GetJobsAllMeRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type GetJobsAllByKeywordsReq struct {
	g.Meta `path:"/all" tags:"Job" method:"get" summary:"Get Jobs All By Keywords"`
}

type GetJobsAllByKeywordsRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type GetJobsByCompanyIDReq struct {
	g.Meta `path:"/company/:company_id" tags:"Job" method:"get" summary:"Get Jobs By Company ID"`
}

type GetJobsByCompanyIDRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type GetJobsMeCompanyReq struct {
	g.Meta `path:"/me/company" tags:"Job" method:"get" summary:"Get Job Me Company"`
}

type GetJobsMeCompanyRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type GetJobsMeUserReq struct {
	g.Meta `path:"/me/user" tags:"Job" method:"get" summary:"Get Job Me User"`
}

type GetJobsMeUserRes struct {
	Jobs map[string][]*entity.Job `json:"jobs"`
}

type GetJobByIDCompanyReq struct {
	g.Meta `path:"/:job_id/company" tags:"Job" method:"get" summary:"Get Job By ID Company"`
}

type GetJobByIDCompanyRes struct {
	Job *entity.Job `json:"job"`
}

type GetJobByIDUserReq struct {
	g.Meta `path:"/:job_id/user" tags:"Job" method:"get" summary:"Get Job By ID User"`
}

type GetJobByIDUserRes struct {
	Job *entity.Job `json:"job"`
}

type GetJobsCreatedByCompanyIDReq struct {
	g.Meta `path:"/company/:company_id/createdby" tags:"Job" method:"get" summary:"Get Job Created By Company ID"`
}

type GetJobsCreatedByCompanyIDRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type GetJobsCreatedByUserIDReq struct {
	g.Meta `path:"/user/:user_id/createdby" tags:"Job" method:"get" summary:"Get Job Created By User ID"`
}

type GetJobsCreatedByUserIDRes struct {
	Jobs []*entity.Job `json:"jobs"`
}

type PatchUpdateJobByIDReq struct {
	g.Meta `path:"/:job_id" tags:"Job" method:"patch" summary:"Update Job By ID"`
	Title         string      `json:"title"            v:"required"`
	Tag           string      `json:"tag"              v:"required"`
	Description	  string      `json:"description"      v:"required"`
	Level	  			string      `json:"level"            v:"required"`
	Salary	  		string      `json:"salary"           v:"required"`
	Location	  	string      `json:"location"         v:"required"`
	IsRemote	  	string      `json:"is_remote"        v:"required"`
	IsHybrid	  	string      `json:"is_hybrid"        v:"required"`
	Expiry	  		gtime.Time  `json:"expiry"           v:"required"`
	Status	  	  string      `json:"status"           v:"required"`
}

type PatchUpdateJobByIDRes struct {
}

type DeleteJobByIDReq struct {
	g.Meta `path:"/:job_id" tags:"Job" method:"delete" summary:"Delete Job By ID"`
}

type DeleteJobByIDRes struct {
}

