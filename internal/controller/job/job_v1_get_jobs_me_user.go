package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetJobsMeUser(ctx context.Context, req *v1.GetJobsMeUserReq) (res *v1.GetJobsMeUserRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {

		return
	}

	jobs, err := service.Job().GetJobsByUserID(ctx, *tokenData.UserID)
	
	res = &v1.GetJobsMeUserRes{
		Jobs: jobs,
	}

	return
}
