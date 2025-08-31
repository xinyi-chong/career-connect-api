package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetJobsAllMe(ctx context.Context, req *v1.GetJobsAllMeReq) (res *v1.GetJobsAllMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.UserID == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
		return
	}
	
	jobs, err := service.Job().GetJobsAll(ctx)

	res = &v1.GetJobsAllMeRes{
		Jobs: jobs,
	}

	return
}
