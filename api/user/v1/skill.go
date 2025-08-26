package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Skills
type GetSkillsReq struct {
	g.Meta `path:"/skills" tags:"User" method:"get" summary:"Get Skills"`
}

type GetSkillsRes struct {
	Skills     []*entity.Skill      `json:"skills"`
}

type GetSkillsByUserIDReq struct {
	g.Meta `path:"/:user_id/skills" tags:"User" method:"get" summary:"Get Skills By User ID"`
	UserID     string      `json:"user_id"     v:"required"`
}

type GetSkillsByUserIDRes struct {
	Skills     []*entity.Skill      `json:"skills"`
}

type PostCreateSkillReq struct {
	g.Meta `path:"/skill" tags:"User" method:"post" summary:"Create Skill"`
	Name     string      `json:"name"     v:"required"`
}

type PostCreateSkillRes struct {
	Id     string      `json:"id"`
}

type PatchUpdateSkillByIDReq struct {
	g.Meta `path:"/skill/:skill_id" tags:"User" method:"patch" summary:"Update Skill By ID"`
	Name     string      `json:"name"     v:"required"`
}

type PatchUpdateSkillByIDRes struct {
}

type DeleteSkillByIDReq struct {
	g.Meta `path:"/skill/:skill_id" tags:"User" method:"delete" summary:"Delete Skill By ID"`
}

type DeleteSkillByIDRes struct {
}