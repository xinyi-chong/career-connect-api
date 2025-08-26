package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteApplicationChatByApplicationID(ctx context.Context, req *v1.DeleteApplicationChatByApplicationIDReq) (res *v1.DeleteApplicationChatByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	err = service.Application().DeleteApplicationChatByApplicationID(ctx, applicationID)

	return
}
