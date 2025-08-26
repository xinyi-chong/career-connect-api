package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateCompanySubscription(ctx context.Context, req *v1.PostCreateCompanySubscriptionReq) (res *v1.PostCreateCompanySubscriptionRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	id, err := service.Company().PostCreateCompanySubscription(ctx, req, *tokenData.CompanyID)
	if err != nil {
		return
	}

	res = &v1.PostCreateCompanySubscriptionRes{
		Id: *id,
	}

	return
}
