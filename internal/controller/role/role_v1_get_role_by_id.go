package role

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetRoleByID(ctx context.Context, req *v1.GetRoleByIDReq) (res *v1.GetRoleByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
 
	role, err := service.Role().GetRoleByID(ctx, roleID)

	res = &v1.GetRoleByIDRes{
		Role: role,
	}

	return
}
