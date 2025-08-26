package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) ForgetPassword(ctx context.Context, req *v1.ForgetPasswordReq) (res *v1.ForgetPasswordRes, err error) {
	err = service.Auth().ForgetPassword(ctx, req)

	return
}
