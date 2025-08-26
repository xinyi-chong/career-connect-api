package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetUserSubscriptionMe(ctx context.Context, req *v1.GetUserSubscriptionMeReq) (res *v1.GetUserSubscriptionMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}

	userSubscription, err := service.User().GetUserSubscriptionByUserID(ctx, *tokenData.UserID)
	
	res = &v1.GetUserSubscriptionMeRes{
		UserSubscription: userSubscription,
	}

	return
}
