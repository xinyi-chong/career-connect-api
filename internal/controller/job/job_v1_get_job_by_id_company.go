package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetJobByIDCompany(ctx context.Context, req *v1.GetJobByIDCompanyReq) (res *v1.GetJobByIDCompanyRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()

	job, err := service.Job().GetJobByID(ctx, jobID)
	if err != nil {
		return
	}

	res = &v1.GetJobByIDCompanyRes{
		Job: job,
	}

	return
}
