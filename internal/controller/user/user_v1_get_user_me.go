package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetUserMe(ctx context.Context, req *v1.GetUserMeReq) (res *v1.GetUserMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	user, err := service.User().GetUserByID(ctx, *tokenData.UserID)

	res = &v1.GetUserMeRes{
		User: *user,
	}

	return
}
