// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package suggest_skill

import (
	"context"

	"gf_demo/api/suggest_skill/v1"
)

type ISuggestSkillV1 interface {
	PostCreateSuggestSkill(ctx context.Context, req *v1.PostCreateSuggestSkillReq) (res *v1.PostCreateSuggestSkillRes, err error)
	GetSuggestSkillByID(ctx context.Context, req *v1.GetSuggestSkillByIDReq) (res *v1.GetSuggestSkillByIDRes, err error)
	PatchUpdateSuggestSkillByID(ctx context.Context, req *v1.PatchUpdateSuggestSkillByIDReq) (res *v1.PatchUpdateSuggestSkillByIDRes, err error)
	DeleteSuggestSkillByID(ctx context.Context, req *v1.DeleteSuggestSkillByIDReq) (res *v1.DeleteSuggestSkillByIDRes, err error)
}
