package permission

import (
	"context"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreatePermission(ctx context.Context, req *v1.PostCreatePermissionReq) (res *v1.PostCreatePermissionRes, err error) {
	permissionID, err := service.Permission().PostCreatePermission(ctx, req)

	res = &v1.PostCreatePermissionRes{
		Id: permissionID,
	}
	
	return
}
