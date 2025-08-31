package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// User Plan
type PostCreateUserPlanReq struct {
	g.Meta `path:"/plan" tags:"User" method:"post" summary:"Create User Plan"`
	Name     string      `json:"name"     v:"required"`
}

type PostCreateUserPlanRes struct {
	Id     string      `json:"id"`
}

type GetUserPlanByIDReq struct {
	g.Meta `path:"/plan/:user_plan_id" tags:"User" method:"get" summary:"Get User Plan By ID"`
}

type GetUserPlanByIDRes struct {
	UserPlan     *entity.UserPlan      `json:"user_plan"`
}

type PatchUpdateUserPlanByIDReq struct {
	g.Meta `path:"/plan/:user_plan_id" tags:"User" method:"patch" summary:"Update User Plan By ID"`
	Name     string      `json:"name"     v:"required"`
}

type PatchUpdateUserPlanByIDRes struct {
}

type DeleteUserPlanByIDReq struct {
	g.Meta `path:"/plan/:user_plan_id" tags:"User" method:"delete" summary:"Delete User Plan By ID"`
}

type DeleteUserPlanByIDRes struct {
}