package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetJobsByCompanyID(ctx context.Context, req *v1.GetJobsByCompanyIDReq) (res *v1.GetJobsByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	jobs, err := service.Job().GetJobsByCompanyID(ctx, companyID)

	res = &v1.GetJobsByCompanyIDRes{
		Jobs: jobs,
	}

	return
}
