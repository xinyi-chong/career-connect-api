package suggest_skill

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/suggest_skill/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetSuggestSkillByID(ctx context.Context, req *v1.GetSuggestSkillByIDReq) (res *v1.GetSuggestSkillByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	suggestSkillID := r.GetRouter("suggest_skill_id").String()
 
	suggestSkill, err := service.SuggestSkill().GetSuggestSkillByID(ctx, suggestSkillID)

	res = &v1.GetSuggestSkillByIDRes{
		SuggestSkill: suggestSkill,
	}

	return
}
