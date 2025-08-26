package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateCompanyPlan(ctx context.Context, req *v1.PostCreateCompanyPlanReq) (res *v1.PostCreateCompanyPlanRes, err error) {
	id, err := service.Company().PostCreateCompanyPlan(ctx, req)
	if err != nil {
		return
	}

	res = &v1.PostCreateCompanyPlanRes{
		Id: *id,
	}

	return
}
