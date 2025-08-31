package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteSkillByID(ctx context.Context, req *v1.DeleteSkillByIDReq) (res *v1.DeleteSkillByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	skillID := r.GetRouter("skill_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().DeleteSkillByID(ctx, skillID, *tokenData.UserID)

	return
}
