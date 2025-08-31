package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteApplicationFileByApplicationIDFileID(ctx context.Context, req *v1.DeleteApplicationFileByApplicationIDFileIDReq) (res *v1.DeleteApplicationFileByApplicationIDFileIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()
	fileID := r.GetRouter("file_id").String()

	err = service.Application().DeleteApplicationFileByApplicationIDFileID(ctx, applicationID, fileID)

	return
}
