package application

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetScheduleByApplicationIDCompany(ctx context.Context, req *v1.GetScheduleByApplicationIDCompanyReq) (res *v1.GetScheduleByApplicationIDCompanyRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.CompanyID == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
		return
	}

	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	schedules, err := service.Application().GetScheduleByApplicationIDCompany(ctx, applicationID, *tokenData.CompanyID)
	
	if err != nil {
		return
	}

	res = &v1.GetScheduleByApplicationIDCompanyRes{
		Schedules: schedules,
	}
	
	return
}
