package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeletePermissionByID(ctx context.Context, req *v1.DeletePermissionByIDReq) (res *v1.DeletePermissionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	permissionID := r.GetRouter("permission_id").String()
 
	err = service.Permission().DeletePermissionByID(ctx, permissionID);

	return
}
