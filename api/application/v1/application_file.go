package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Files
type PostCreateApplicationFilesByApplicationIDReq struct {
	g.Meta `path:"/:application_id/file" tags:"Application" method:"post" summary:"Create Application Files By Application ID"`
	Resumes				 []*ghttp.UploadFile      `json:"resumes"`
	OtherFiles		 []*ghttp.UploadFile      `json:"other_files"`
}

type PostCreateApplicationFilesByApplicationIDRes struct {
	SuccessUploadedResumes      int  `json:"success_uploaded_resumes"`
	SuccessUploadedOtherFiles   int  `json:"success_uploaded_other_files"`
}

type PostCreateApplicationFileByApplicationIDResumeIDReq struct {
	g.Meta `path:"/:application_id/file/resume/:resume_id" tags:"Application" method:"post" summary:"Create Application File By Application ID & Resume ID"`
}

type PostCreateApplicationFileByApplicationIDResumeIDRes struct {
	Id   *string  `json:"id"`
}

type GetApplicationFilesByApplicationIDReq struct {
	g.Meta `path:"/:application_id/files" tags:"Application" method:"get" summary:"Get Application Files By Application ID"`
}

type GetApplicationFilesByApplicationIDRes struct {
	Resumes				 []*entity.ApplicationFile      `json:"resumes"`
	OtherFiles		 []*entity.ApplicationFile      `json:"other_files"`
}

type DeleteApplicationFileByApplicationIDFileIDReq struct {
	g.Meta `path:"/:application_id/file/:file_id" tags:"Application" method:"delete" summary:"Delete Application File By Application ID & File ID"`
}

type DeleteApplicationFileByApplicationIDFileIDRes struct {
}