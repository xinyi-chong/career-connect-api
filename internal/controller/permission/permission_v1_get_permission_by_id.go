package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetPermissionByID(ctx context.Context, req *v1.GetPermissionByIDReq) (res *v1.GetPermissionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	permissionID := r.GetRouter("permission_id").String()
 
	permission, err := service.Permission().GetPermissionByID(ctx, permissionID);
	
	res = &v1.GetPermissionByIDRes{
		Permission: permission,
	}

	return
}
