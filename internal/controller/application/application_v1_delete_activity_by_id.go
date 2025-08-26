package application

import (
	"context"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteActivityByID(ctx context.Context, req *v1.DeleteActivityByIDReq) (res *v1.DeleteActivityByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	activityID := r.GetRouter("activity_id").String()

	err = service.Application().DeleteActivityByID(ctx, activityID)

	return
}
