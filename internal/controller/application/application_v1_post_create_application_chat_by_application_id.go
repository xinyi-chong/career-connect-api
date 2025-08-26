package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateApplicationChatByApplicationID(ctx context.Context, req *v1.PostCreateApplicationChatByApplicationIDReq) (res *v1.PostCreateApplicationChatByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	applicationChatID, err := service.Application().PostCreateApplicationChatByApplicationID(ctx, req, applicationID)

	if err != nil {
		return
	}

	res = &v1.PostCreateApplicationChatByApplicationIDRes{
		Id: applicationChatID,
	}

	return
}
