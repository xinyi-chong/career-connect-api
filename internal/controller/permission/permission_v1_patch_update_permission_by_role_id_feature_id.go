package permission

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdatePermissionByRoleIDFeatureID(ctx context.Context, req *v1.PatchUpdatePermissionByRoleIDFeatureIDReq) (res *v1.PatchUpdatePermissionByRoleIDFeatureIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
	featureID := r.GetRouter("feature_id").String()
 
	err = service.Permission().PatchUpdatePermissionByRoleIDFeatureID(ctx, req, roleID, featureID);

	return
}
