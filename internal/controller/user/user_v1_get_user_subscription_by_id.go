package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetUserSubscriptionByID(ctx context.Context, req *v1.GetUserSubscriptionByIDReq) (res *v1.GetUserSubscriptionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	subscriptionID := r.GetRouter("subscription_id").String()

	userSubscription, err := service.User().GetUserSubscriptionByID(ctx, subscriptionID)
	
	res = &v1.GetUserSubscriptionByIDRes{
		UserSubscription: userSubscription,
	}

	return
}
