package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetEducationsByUserID(ctx context.Context, req *v1.GetEducationsByUserIDReq) (res *v1.GetEducationsByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	educations, err := service.User().GetEducationsByUserID(ctx, userID);

	res = &v1.GetEducationsByUserIDRes{
		Educations: educations,
	}

	return
}
