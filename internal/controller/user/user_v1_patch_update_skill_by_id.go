package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateSkillByID(ctx context.Context, req *v1.PatchUpdateSkillByIDReq) (res *v1.PatchUpdateSkillByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	skillID := r.GetRouter("skill_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().PatchUpdateSkillByID(ctx, req, skillID, *tokenData.UserID)

	return
}
