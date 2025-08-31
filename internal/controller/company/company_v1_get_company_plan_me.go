package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCompanyPlanMe(ctx context.Context, req *v1.GetCompanyPlanMeReq) (res *v1.GetCompanyPlanMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	companySubscription, err := service.Company().GetCompanySubscriptionByCompanyID(ctx, *tokenData.CompanyID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyPlanMeRes{
		CompanyPlan: companySubscription.CompanyPlan,
	}

	return
}
