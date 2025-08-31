package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Application Activity
type PostCreateActivityReq struct {
	g.Meta `path:"/activity" tags:"Application" method:"post" summary:"Create Activity"`
	Name        string      `json:"name"       v:"required"`
}

type PostCreateActivityRes struct {
	Id					*string  		`json:"id"`
}

type PostCreateActivityByApplicationIDReq struct {
	g.Meta `path:"/:application_id/activity" tags:"Application" method:"post" summary:"Create Activity By Application ID"`
	Name        string      `json:"name"       v:"required"`
}

type PostCreateActivityByApplicationIDRes struct {
	Id					*string  		`json:"id"`
}

type GetActivityByIDReq struct {
	g.Meta `path:"/activity/:activity_id" tags:"Application" method:"get" summary:"Get Activity By Application ID"`
}

type GetActivityByIDRes struct {
	Activity *entity.Activity `json:"activity"`
}

type PatchUpdateActivityByIDReq struct {
	g.Meta `path:"/activity/:activity_id" tags:"Application" method:"patch" summary:"Update Activity By ID"`
	Name        string      `json:"name"       v:"required"`
}

type PatchUpdateActivityByIDRes struct {
}


type DeleteActivityByIDReq struct {
	g.Meta `path:"/activity/:activity_id" tags:"Application" method:"delete" summary:"Delete Activity By ID"`
}

type DeleteActivityByIDRes struct {
}