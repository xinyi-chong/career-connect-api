package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteApplicationByID(ctx context.Context, req *v1.DeleteApplicationByIDReq) (res *v1.DeleteApplicationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	err = service.Application().DeleteApplicationByID(ctx, applicationID)
	
	return
}
