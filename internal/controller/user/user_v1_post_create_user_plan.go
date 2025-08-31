package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateUserPlan(ctx context.Context, req *v1.PostCreateUserPlanReq) (res *v1.PostCreateUserPlanRes, err error) {
	id, err := service.User().PostCreateUserPlan(ctx, req)

	res = &v1.PostCreateUserPlanRes{
		Id: *id,
	}

	return
}
