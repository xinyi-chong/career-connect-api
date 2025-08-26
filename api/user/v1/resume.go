package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Resumes
type GetResumesReq struct {
	g.Meta `path:"/resumes" tags:"User" method:"get" summary:"Get Resumes"`
}

type GetResumesRes struct {
	Resumes     []*entity.Resume      `json:"resumes"`
}

type GetResumeByIDReq struct {
	g.Meta `path:"/resume/:resume_id" tags:"User" method:"get" summary:"Get Resume By ID"`
}

type GetResumeByIDRes struct {
	Resume     *entity.Resume      `json:"resume"`
}

type PostCreateResumeReq struct {
	g.Meta `path:"/resume" tags:"User" method:"post" summary:"Create Resume"`
	Resume     *ghttp.UploadFile   `json:"resume"     v:"required"`
}

type PostCreateResumeRes struct {
	Id     string      `json:"id"`
}

type PatchUpdateResumeByIDReq struct {
	g.Meta `path:"/resume/:resume_id" tags:"User" method:"patch" summary:"Update Resume By ID"`
	Resume     *ghttp.UploadFile      `json:"resume"     v:"required"`
}

type PatchUpdateResumeByIDRes struct {
}

type DeleteResumeByIDReq struct {
	g.Meta `path:"/resume/:resume_id" tags:"User" method:"delete" summary:"Delete Resume By ID"`
}

type DeleteResumeByIDRes struct {
}