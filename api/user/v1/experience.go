package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Experience
type PostCreateExperienceReq struct {
	g.Meta `path:"/experience" tags:"User" method:"post" summary:"Create Experience"`
	StartDate     gtime.Time   `json:"start_date"     v:"required"`
	EndDate     	*gtime.Time  `json:"end_date"`
	Description   *string      `json:"description"`
	Title     		string       `json:"title"          v:"required"`
	CompanyID     *string      `json:"company_id"`
	CompanyString *string      `json:"company_string"`
}

type PostCreateExperienceRes struct {
	Id     string      `json:"id"`
}

type GetExperiencesReq struct {
	g.Meta `path:"/experiences" tags:"User" method:"get" summary:"Get Experiences"`
}

type GetExperiencesRes struct {
	Experiences     []*entity.Experience      `json:"experiences"`
}

type GetExperienceByIDReq struct {
	g.Meta `path:"/experience/:experience_id" tags:"User" method:"get" summary:"Get Experience By ID"`
}

type GetExperienceByIDRes struct {
	Experience     *entity.Experience      `json:"experience"`
}

type GetExperiencesByUserIDReq struct {
	g.Meta `path:"/:user_id/experiences" tags:"User" method:"get" summary:"Get Experiences By User ID"`
	UserID     string      `json:"user_id"     v:"required"`
}

type GetExperiencesByUserIDRes struct {
	Experiences     []*entity.Experience      `json:"experiences"`
}

type PatchUpdateExperienceByIDReq struct {
	g.Meta `path:"/experience/:experience_id" tags:"User" method:"patch" summary:"Update Experience By ID"`
	StartDate     *gtime.Time   `json:"start_date"`
	EndDate     	*gtime.Time  `json:"end_date"`
	IsPresent     *bool      	 `json:"is_present"`
	Description   *string      `json:"description"`
	Title     		*string       `json:"title"`
	CompanyID     *string      `json:"company_id"`
	CompanyString *string      `json:"company_string"`
}

type PatchUpdateExperienceByIDRes struct {
}

type DeleteExperienceByIDReq struct {
	g.Meta `path:"/experience/:experience_id" tags:"User" method:"delete" summary:"Delete Experience By ID"`
}

type DeleteExperienceByIDRes struct {
}