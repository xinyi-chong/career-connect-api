package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteUserSubscriptionMe(ctx context.Context, req *v1.DeleteUserSubscriptionMeReq) (res *v1.DeleteUserSubscriptionMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	err = service.User().DeleteUserSubscriptionByUserID(ctx, *tokenData.UserID)

	return
}
