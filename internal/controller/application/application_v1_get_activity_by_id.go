package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetActivityByID(ctx context.Context, req *v1.GetActivityByIDReq) (res *v1.GetActivityByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	activityID := r.GetRouter("activity_id").String()

	activity, err := service.Application().GetActivityByID(ctx, activityID)

	if err != nil {
		return
	}

	res = &v1.GetActivityByIDRes{
		Activity: activity,
	}

	return
}
