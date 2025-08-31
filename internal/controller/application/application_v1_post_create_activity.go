package application

import (
	"context"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateActivity(ctx context.Context, req *v1.PostCreateActivityReq) (res *v1.PostCreateActivityRes, err error) {
	activityID, err := service.Application().PostCreateActivity(ctx, req)

	if err != nil {
		return
	}

	res = &v1.PostCreateActivityRes{
		Id: activityID,
	}

	return
}
