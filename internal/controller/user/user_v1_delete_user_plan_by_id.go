package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteUserPlanByID(ctx context.Context, req *v1.DeleteUserPlanByIDReq) (res *v1.DeleteUserPlanByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userPlanID := r.GetRouter("user_plan_id").String()

	err = service.User().DeleteUserPlanByID(ctx, userPlanID)

	return
}
