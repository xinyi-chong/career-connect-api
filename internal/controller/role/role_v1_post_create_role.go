package role

import (
	"context"

	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateRole(ctx context.Context, req *v1.PostCreateRoleReq) (res *v1.PostCreateRoleRes, err error) {
	roleID, err := service.Role().PostCreateRole(ctx, req)

	res = &v1.PostCreateRoleRes{
		Id: roleID,
	}

	return
}
