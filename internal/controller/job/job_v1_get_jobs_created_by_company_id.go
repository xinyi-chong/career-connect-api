package job

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetJobsCreatedByCompanyID(ctx context.Context, req *v1.GetJobsCreatedByCompanyIDReq) (res *v1.GetJobsCreatedByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	jobs, err := service.Job().GetJobsCreatedByCompanyID(ctx, companyID)

	res = &v1.GetJobsCreatedByCompanyIDRes{
		Jobs: jobs,
	}

	return
}
