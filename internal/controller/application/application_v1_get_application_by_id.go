package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetApplicationByID(ctx context.Context, req *v1.GetApplicationByIDReq) (res *v1.GetApplicationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	application, err := service.Application().GetApplicationByID(ctx, applicationID)

	if err != nil {
		return
	}

	res = &v1.GetApplicationByIDRes{
		Application: application,
	}

	return
}
