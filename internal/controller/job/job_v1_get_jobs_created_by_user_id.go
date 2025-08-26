package job

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetJobsCreatedByUserID(ctx context.Context, req *v1.GetJobsCreatedByUserIDReq) (res *v1.GetJobsCreatedByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	jobs, err := service.Job().GetJobsCreatedByUserID(ctx, userID)

	res = &v1.GetJobsCreatedByUserIDRes{
		Jobs: jobs,
	}

	return
}
