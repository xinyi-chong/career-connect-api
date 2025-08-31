package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	accessToken, err := service.Token().RefreshToken(ctx, req.RefreshToken)
	
	if err != nil {
		return
	}
	
	res = &v1.RefreshTokenRes{
		AccessToken: accessToken,
	}
	
	return
}
