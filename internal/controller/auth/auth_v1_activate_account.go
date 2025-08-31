package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) ActivateAccount(ctx context.Context, req *v1.ActivateAccountReq) (res *v1.ActivateAccountRes, err error) {
	err = service.Auth().ActivateAccount(ctx, req)
	
	return
}
