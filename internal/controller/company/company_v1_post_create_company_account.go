package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateCompanyAccount(ctx context.Context, req *v1.PostCreateCompanyAccountReq) (res *v1.PostCreateCompanyAccountRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	in := &v1.PostCreateCompanyAccountByCompanyIDReq{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	err = service.Company().PostCreateCompanyAccountByCompanyID(ctx, in, *tokenData.CompanyID)

	return
}
