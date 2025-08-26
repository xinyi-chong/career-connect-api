package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateScheduleByApplicationID(ctx context.Context, req *v1.PostCreateScheduleByApplicationIDReq) (res *v1.PostCreateScheduleByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	scheduleID, err := service.Application().PostCreateScheduleByApplicationID(ctx, req, applicationID)
	
	if err != nil {
		return
	}
	
	res = &v1.PostCreateScheduleByApplicationIDRes{
		Id: scheduleID,
	}
	
	return
}