package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) SignInUser(ctx context.Context, req *v1.SignInUserReq) (res *v1.SignInUserRes, err error) {
	user, err := service.Auth().SignInUser(ctx, req)
	if err != nil {
		return
	}

	accessToken, refreshToken, err := service.Token().GenerateAccessAndRefreshToken(ctx, user.AccountId)
	if err != nil {
		return
	}

	res = &v1.SignInUserRes{
		User: user,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
	
	return
}
