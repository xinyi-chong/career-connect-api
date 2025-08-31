package job

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteJobByID(ctx context.Context, req *v1.DeleteJobByIDReq) (res *v1.DeleteJobByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()

	err = service.Job().DeleteJobByID(ctx, jobID)

	return
}
