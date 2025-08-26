package suggest_skill

import (
	"context"

	v1 "gf_demo/api/suggest_skill/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateSuggestSkill(ctx context.Context, req *v1.PostCreateSuggestSkillReq) (res *v1.PostCreateSuggestSkillRes, err error) {
	suggestSkillID, err := service.SuggestSkill().PostCreateSuggestSkill(ctx, req)

	res = &v1.PostCreateSuggestSkillRes{
		Id: suggestSkillID,
	}
	
	return
}
