package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetJobByIDUser(ctx context.Context, req *v1.GetJobByIDUserReq) (res *v1.GetJobByIDUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()

	job, err := service.Job().GetJobByID(ctx, jobID)

	res = &v1.GetJobByIDUserRes{
		Job: job,
	}

	return
}
