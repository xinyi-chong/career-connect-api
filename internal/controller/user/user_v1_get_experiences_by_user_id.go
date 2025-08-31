package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetExperiencesByUserID(ctx context.Context, req *v1.GetExperiencesByUserIDReq) (res *v1.GetExperiencesByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	experiences, err := service.User().GetExperiencesByUserID(ctx, userID);

	res = &v1.GetExperiencesByUserIDRes{
		Experiences: experiences,
	}

	return
}
