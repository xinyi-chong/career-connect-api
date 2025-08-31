package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCompanyAccountsMe(ctx context.Context, req *v1.GetCompanyAccountsMeReq) (res *v1.GetCompanyAccountsMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}

	companyAccounts, err := service.Company().GetCompanyAccountsByCompanyID(ctx, *tokenData.CompanyID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyAccountsMeRes{
		CompanyAccounts: companyAccounts,
	}

	return
}
