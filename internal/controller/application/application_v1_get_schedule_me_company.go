package application

import (
	"context"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetScheduleMeCompany(ctx context.Context, req *v1.GetScheduleMeCompanyReq) (res *v1.GetScheduleMeCompanyRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.CompanyID == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
		return
	}

	schedules, err := service.Application().GetScheduleByCompanyID(ctx, *tokenData.CompanyID)
	
	if err != nil {
		return
	}

	res = &v1.GetScheduleMeCompanyRes{
		Schedules: schedules,
	}

	return
}
