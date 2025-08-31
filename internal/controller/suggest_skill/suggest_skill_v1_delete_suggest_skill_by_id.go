package suggest_skill

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/suggest_skill/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteSuggestSkillByID(ctx context.Context, req *v1.DeleteSuggestSkillByIDReq) (res *v1.DeleteSuggestSkillByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	suggestSkillID := r.GetRouter("suggest_skill_id").String()
 
	err = service.SuggestSkill().DeleteSuggestSkillByID(ctx, suggestSkillID)
	
	return
}
