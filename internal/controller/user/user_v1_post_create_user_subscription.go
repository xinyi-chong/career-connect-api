package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateUserSubscription(ctx context.Context, req *v1.PostCreateUserSubscriptionReq) (res *v1.PostCreateUserSubscriptionRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	id, err := service.User().PostCreateUserSubscription(ctx, req, *tokenData.UserID)

	res = &v1.PostCreateUserSubscriptionRes{
		Id: *id,
	}

	return
}
