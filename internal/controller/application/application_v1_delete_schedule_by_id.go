package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteScheduleByID(ctx context.Context, req *v1.DeleteScheduleByIDReq) (res *v1.DeleteScheduleByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	scheduleID := r.GetRouter("schedule_id").String()

	err = service.Application().DeleteScheduleByID(ctx, scheduleID)

	return
}
