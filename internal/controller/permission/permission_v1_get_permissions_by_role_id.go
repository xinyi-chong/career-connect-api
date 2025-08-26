package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetPermissionsByRoleID(ctx context.Context, req *v1.GetPermissionsByRoleIDReq) (res *v1.GetPermissionsByRoleIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
 
	permissions, err := service.Permission().GetPermissionsByRoleID(ctx, roleID);

	res = &v1.GetPermissionsByRoleIDRes{
		Permissions: permissions,
	}

	return
}
