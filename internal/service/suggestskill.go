// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/suggest_skill/v1"
	"gf_demo/internal/model/entity"
)

type (
	ISuggestSkill interface {
		GetSuggestSkillByID(ctx context.Context, suggestSkillID string) (*entity.SuggestSkill, error)
		PostCreateSuggestSkill(ctx context.Context, req *v1.PostCreateSuggestSkillReq) (*string, error)
		PatchUpdateSuggestSkillByID(ctx context.Context, req *v1.PatchUpdateSuggestSkillByIDReq, suggestSkillID string) error
		DeleteSuggestSkillByID(ctx context.Context, suggestSkillID string) error
	}
)

var (
	localSuggestSkill ISuggestSkill
)

func SuggestSkill() ISuggestSkill {
	if localSuggestSkill == nil {
		panic("implement not found for interface ISuggestSkill, forgot register?")
	}
	return localSuggestSkill
}

func RegisterSuggestSkill(i ISuggestSkill) {
	localSuggestSkill = i
}
