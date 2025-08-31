package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateCompanySubscriptionByID(ctx context.Context, req *v1.PatchUpdateCompanySubscriptionByIDReq) (res *v1.PatchUpdateCompanySubscriptionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	subscriptionID := r.GetRouter("subscription_id").String()

	err = service.Company().PatchUpdateCompanySubscriptionByID(ctx, req, subscriptionID)

	return
}
