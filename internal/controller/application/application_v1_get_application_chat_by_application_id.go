package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetApplicationChatByApplicationID(ctx context.Context, req *v1.GetApplicationChatByApplicationIDReq) (res *v1.GetApplicationChatByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	applicationChats, err := service.Application().GetApplicationChatByApplicationID(ctx, applicationID)

	if err != nil {
		return
	}

	res = &v1.GetApplicationChatByApplicationIDRes{
		ApplicationChats: applicationChats,
	}

	return
}
