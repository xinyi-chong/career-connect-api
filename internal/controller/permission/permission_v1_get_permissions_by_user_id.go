package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCompanyPermissionsByUserID(ctx context.Context, req *v1.GetCompanyPermissionsByUserIDReq) (res *v1.GetCompanyPermissionsByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()
 
	permissions, err := service.Permission().GetCompanyPermissionsByUserID(ctx, userID);
	
	res = &v1.GetCompanyPermissionsByUserIDRes{
		Permissions: permissions,
	}

	return
}
