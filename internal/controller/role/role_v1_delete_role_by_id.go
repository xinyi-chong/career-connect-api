package role

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteRoleByID(ctx context.Context, req *v1.DeleteRoleByIDReq) (res *v1.DeleteRoleByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	roleID := r.GetRouter("role_id").String()
 
	err = service.Role().DeleteRoleByID(ctx, roleID)
	
	return
}
