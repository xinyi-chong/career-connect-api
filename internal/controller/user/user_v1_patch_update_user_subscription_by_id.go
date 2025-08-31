package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateUserSubscriptionByID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionByIDReq) (res *v1.PatchUpdateUserSubscriptionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	subscriptionID := r.GetRouter("subscription_id").String()

	err = service.User().PatchUpdateUserSubscriptionByID(ctx, req, subscriptionID)
	
	return
}
