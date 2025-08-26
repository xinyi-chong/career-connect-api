package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateApplicationByID(ctx context.Context, req *v1.PatchUpdateApplicationByIDReq) (res *v1.PatchUpdateApplicationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	err = service.Application().PatchUpdateApplicationByID(ctx, req, applicationID)

	return
}
