package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) SignInCompany(ctx context.Context, req *v1.SignInCompanyReq) (res *v1.SignInCompanyRes, err error) {
	company, err := service.Auth().SignInCompany(ctx, req)
	if err != nil {
		return
	}

	accessToken, refreshToken, err := service.Token().GenerateAccessAndRefreshToken(ctx, company.AccountId)
	if err != nil {
		return
	}

	res = &v1.SignInCompanyRes{
		Company: company,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	
	return
}
