package suggest_skill

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/suggest_skill/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateSuggestSkillByID(ctx context.Context, req *v1.PatchUpdateSuggestSkillByIDReq) (res *v1.PatchUpdateSuggestSkillByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	suggestSkillID := r.GetRouter("suggest_skill_id").String()
 
	err = service.SuggestSkill().PatchUpdateSuggestSkillByID(ctx, req, suggestSkillID)

	return
}
