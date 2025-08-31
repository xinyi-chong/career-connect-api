package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetJobsMeCompany(ctx context.Context, req *v1.GetJobsMeCompanyReq) (res *v1.GetJobsMeCompanyRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {

		return
	}
	
	jobs, err := service.Job().GetJobsByCompanyID(ctx, *tokenData.CompanyID)
	
	res = &v1.GetJobsMeCompanyRes{
		Jobs: jobs,
	}

	return
}
