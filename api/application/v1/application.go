package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type PostCreateApplicationByJobIDReq struct {
	g.Meta `path:"/job/:job_id" tags:"Application" method:"post" summary:"Create Application By Job ID"`
	Answer    	 *string             `json:"answer"`
	ActivityID   string              `json:"activity_id"  v:"required"`
	Resumes      []*ghttp.UploadFile `json:"resumes"`
	OtherFiles   []*ghttp.UploadFile `json:"other_files"`
}

type PostCreateApplicationByJobIDRes struct {
	Id					*string  		`json:"id"`
	SuccessUploadedResumes      int  `json:"success_uploaded_resumes"`
	SuccessUploadedOtherFiles   int  `json:"success_uploaded_other_files"`
}

type GetApplicationByJobIDReq struct {
	g.Meta `path:"/job/:job_id" tags:"Application" method:"get" summary:"Get Application By Job ID"`
}

type GetApplicationByJobIDRes struct {
	Applications []*entity.Application `json:"applications"`
}

type GetApplicationsMeReq struct {
	g.Meta `path:"/me" tags:"Application" method:"get" summary:"Get Applications Me"`
}

type GetApplicationsMeRes struct {
	Applications []*entity.Application `json:"applications"`
}

type GetApplicationByIDReq struct {
	g.Meta `path:"/:application_id" tags:"Application" method:"get" summary:"Get Application By ID"`
}

type GetApplicationByIDRes struct {
	Application *entity.Application `json:"Application"`
}

type PatchUpdateApplicationByIDReq struct {
	g.Meta `path:"/:application_id" tags:"Application" method:"patch" summary:"Update Application By ID"`
	Answer    	 *string      `json:"answer"`
	// ResumeID     *string      `json:"resume_id"`
	ActivityID   *string      `json:"activity_id"`
}

type PatchUpdateApplicationByIDRes struct {}

type DeleteApplicationByIDReq struct {
	g.Meta `path:"/:application_id" tags:"Application" method:"delete" summary:"Delete Application By ID"`
}

type DeleteApplicationByIDRes struct {
}