package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateScheduleByID(ctx context.Context, req *v1.PatchUpdateScheduleByIDReq) (res *v1.PatchUpdateScheduleByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	scheduleID := r.GetRouter("schedule_id").String()

	err = service.Application().PatchUpdateScheduleByID(ctx, req, scheduleID)
	
	return
}
