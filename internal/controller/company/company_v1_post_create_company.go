package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateCompany(ctx context.Context, req *v1.PostCreateCompanyReq) (res *v1.PostCreateCompanyRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}

	id, err := service.Company().PostCreateCompany(ctx, req, tokenData.AccountID)
	if err != nil {
		return
	}

	res = &v1.PostCreateCompanyRes{
		Id: *id,
	}

	return
}
