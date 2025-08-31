package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetScheduleByApplicationID(ctx context.Context, req *v1.GetScheduleByApplicationIDReq) (res *v1.GetScheduleByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	schedules, err := service.Application().GetScheduleByApplicationID(ctx, applicationID)
	
	if err != nil {
		return
	}

	res = &v1.GetScheduleByApplicationIDRes{
		Schedules: schedules,
	}
	
	return
}
