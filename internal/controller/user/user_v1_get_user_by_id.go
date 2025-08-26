package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetUserByID(ctx context.Context, req *v1.GetUserByIDReq) (res *v1.GetUserByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	user, err := service.User().GetUserByID(ctx, userID)

	res = &v1.GetUserByIDRes{
		User: user,
	}

	return
}
