package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateActivityByApplicationID(ctx context.Context, req *v1.PostCreateActivityByApplicationIDReq) (res *v1.PostCreateActivityByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	activityID, err := service.Application().PostCreateActivityByApplicationID(ctx, req, applicationID)

	if err != nil {
		return
	}

	res = &v1.PostCreateActivityByApplicationIDRes{
		Id: activityID,
	}

	return
}
