package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetPermissionByRoleIDFeatureID(ctx context.Context, req *v1.GetPermissionByRoleIDFeatureIDReq) (res *v1.GetPermissionByRoleIDFeatureIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
	featureID := r.GetRouter("feature_id").String()
 
	permission, err := service.Permission().GetPermissionByRoleIDFeatureID(ctx, roleID, featureID);

	res = &v1.GetPermissionByRoleIDFeatureIDRes{
		Permission: permission,
	}

	return
}
