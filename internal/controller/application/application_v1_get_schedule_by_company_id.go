package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetScheduleByCompanyID(ctx context.Context, req *v1.GetScheduleByCompanyIDReq) (res *v1.GetScheduleByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	schedules, err := service.Application().GetScheduleByCompanyID(ctx, companyID)
	
	if err != nil {
		return
	}

	res = &v1.GetScheduleByCompanyIDRes{
		Schedules: schedules,
	}
	
	return
}
