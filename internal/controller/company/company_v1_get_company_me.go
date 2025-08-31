package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCompanyMe(ctx context.Context, req *v1.GetCompanyMeReq) (res *v1.GetCompanyMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}
	
	company, err := service.Company().GetCompanyByID(ctx, *tokenData.CompanyID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyMeRes{
		Company: company,
	}

	return
}
