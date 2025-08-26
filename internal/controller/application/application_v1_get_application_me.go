package application

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetApplicationsMe(ctx context.Context, req *v1.GetApplicationsMeReq) (res *v1.GetApplicationsMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.UserID == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
		return
	}

	applications, err := service.Application().GetApplicationsByUserID(ctx, *tokenData.UserID)

	if err != nil {
		return
	}

	res = &v1.GetApplicationsMeRes{
		Applications: applications,
	}

	return
}
