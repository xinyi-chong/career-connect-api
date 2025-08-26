package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

// Testing
func (c *ControllerV1) RegisterUser(ctx context.Context, req *v1.RegisterUserReq) (res *v1.RegisterUserRes, err error) {
	id, err := service.Auth().RegisterUser(ctx, req)
	
	if err != nil {
		return
	}
	
	res = &v1.RegisterUserRes{
		ID: id,
	}

	return
}
