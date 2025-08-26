package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Education
type GetEducationsReq struct {
	g.Meta `path:"/educations" tags:"User" method:"get" summary:"Get Educations"`
}

type GetEducationsRes struct {
	Educations     []*entity.Education      `json:"educations"`
}

type GetEducationsByUserIDReq struct {
	g.Meta `path:"/:user_id/educations" tags:"User" method:"get" summary:"Get Educations By User ID"`
}

type GetEducationsByUserIDRes struct {
	Educations     []*entity.Education      `json:"educations"`
}

type GetEducationByIDReq struct {
	g.Meta `path:"/education/:education_id" tags:"User" method:"get" summary:"Get Education By ID"`
}

type GetEducationByIDRes struct {
	Education *entity.Education `json:"education"`
}

type PostCreateEducationReq struct {
	g.Meta `path:"/education" tags:"User" method:"post" summary:"Create Education"`
	StartDate        gtime.Time    `json:"start_date"        v:"required"`	//Example Format: 2012-12-31 22:00:00
	EndDate     	   *gtime.Time   `json:"end_date"`
	InstituteID      *string       `json:"institute_id"`
	InstituteString  *string       `json:"institute_string"`
	Level     			 string        `json:"level"             v:"required"`
	Programme        string        `json:"programme"         v:"required"`
	Description      *string       `json:"description"`
}

type PostCreateEducationRes struct {
	Id     string      `json:"id"`
}

type PatchUpdateEducationByIDReq struct {
	g.Meta `path:"/education/:education_id" tags:"User" method:"patch" summary:"Update Education By ID"`
	StartDate        *gtime.Time   `json:"start_date"`
	EndDate     	   *gtime.Time   `json:"end_date"`
	InstituteID      *string       `json:"institute_id"`
	InstituteString  *string       `json:"institute_string"`
	Level     			 *string       `json:"level"`
	Programme        *string       `json:"programme"`
	Description      *string       `json:"description"`
}

type PatchUpdateEducationByIDRes struct {
}

type DeleteEducationByIDReq struct {
	g.Meta `path:"/education/:education_id" tags:"User" method:"delete" summary:"Delete Education By ID"`
}

type DeleteEducationByIDRes struct {
}