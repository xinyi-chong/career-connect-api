package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteUserSubscriptionByID(ctx context.Context, req *v1.DeleteUserSubscriptionByIDReq) (res *v1.DeleteUserSubscriptionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userSubscriptionID := r.GetRouter("subscription_id").String()

	err = service.User().DeleteUserSubscriptionByID(ctx, userSubscriptionID)

	return
}
