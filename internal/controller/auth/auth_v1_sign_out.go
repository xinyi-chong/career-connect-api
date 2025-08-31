package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.Session().RemoveSession(ctx)

	return
}
