package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteUserMe(ctx context.Context, req *v1.DeleteUserMeReq) (res *v1.DeleteUserMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().DeleteUserByID(ctx, *tokenData.UserID)

	return
}
