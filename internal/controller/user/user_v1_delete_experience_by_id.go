package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteExperienceByID(ctx context.Context, req *v1.DeleteExperienceByIDReq) (res *v1.DeleteExperienceByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	experienceID := r.GetRouter("experience_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().DeleteExperienceByID(ctx, experienceID, *tokenData.UserID);

	return
}
