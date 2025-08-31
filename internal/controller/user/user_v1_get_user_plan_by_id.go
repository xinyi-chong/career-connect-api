package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetUserPlanByID(ctx context.Context, req *v1.GetUserPlanByIDReq) (res *v1.GetUserPlanByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userPlanID := r.GetRouter("user_plan_id").String()

	userPlan, err := service.User().GetUserPlanByID(ctx, userPlanID)

	res = &v1.GetUserPlanByIDRes{
		UserPlan: userPlan,
	}

	return
}
