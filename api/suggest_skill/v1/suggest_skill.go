package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateSuggestSkillReq struct {
	g.Meta   `path:"/" method:"post" tags:"Suggest Skill" summary:"Create Suggest Skill"`
	Name 			string   `json:"name"       v:"required"`
	Category 	string   `json:"category"   v:"required"`
}

type PostCreateSuggestSkillRes struct{
	Id 				*string				`json:"id"`
}

type GetSuggestSkillByIDReq struct {
	g.Meta   `path:"/:suggest_skill_id" method:"get" tags:"Suggest Skill" summary:"Get Suggest Skill By ID"`
}

type GetSuggestSkillByIDRes struct {
	SuggestSkill *entity.SuggestSkill	`json:"suggest_skill"`
}

type PatchUpdateSuggestSkillByIDReq struct {
	g.Meta   `path:"/:suggest_skill_id" method:"patch" tags:"Suggest Skill" summary:"Update Suggest Skill By ID"`
	Name 			string   `json:"name"       v:"required"`
	Category 	string   `json:"category"   v:"required"`
}
type PatchUpdateSuggestSkillByIDRes struct {}

type DeleteSuggestSkillByIDReq struct {
	g.Meta   `path:"/:suggest_skill_id" method:"delete" tags:"Suggest Skill" summary:"Delete Suggest Skill By ID"`
}

type DeleteSuggestSkillByIDRes struct {}
