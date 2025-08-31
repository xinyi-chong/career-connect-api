package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetApplicationByJobID(ctx context.Context, req *v1.GetApplicationByJobIDReq) (res *v1.GetApplicationByJobIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()

	applications, err := service.Application().GetApplicationsByJobID(ctx, jobID)
	
	if err != nil {
		return
	}

	res = &v1.GetApplicationByJobIDRes{
		Applications: applications,
	}

	return
}
