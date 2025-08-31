package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

// Testing
func (c *ControllerV1) RegisterCompany(ctx context.Context, req *v1.RegisterCompanyReq) (res *v1.RegisterCompanyRes, err error) {
	id, err := service.Auth().RegisterCompany(ctx, req)
	
	if err != nil {
		return
	}

	res = &v1.RegisterCompanyRes{
		ID: id,
	}

	return
}
