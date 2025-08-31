package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetJobsAllByKeywords(ctx context.Context, req *v1.GetJobsAllByKeywordsReq) (res *v1.GetJobsAllByKeywordsRes, err error) {
	keywords := g.RequestFromCtx(ctx).Get("keywords")

	jobs, err := service.Job().GetJobsAllByKeywords(ctx, keywords.Array())
	
	res = &v1.GetJobsAllByKeywordsRes{
		Jobs: jobs,
	}
	
	return
}
