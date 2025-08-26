package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateActivityByID(ctx context.Context, req *v1.PatchUpdateActivityByIDReq) (res *v1.PatchUpdateActivityByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	activityID := r.GetRouter("activity_id").String()

	err = service.Application().PatchUpdateActivityByID(ctx, req, activityID)

	return
}
