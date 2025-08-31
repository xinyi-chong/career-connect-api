package role

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateRoleByID(ctx context.Context, req *v1.PatchUpdateRoleByIDReq) (res *v1.PatchUpdateRoleByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
 
	err = service.Role().PatchUpdateRoleByID(ctx, req, roleID)

	return
}
