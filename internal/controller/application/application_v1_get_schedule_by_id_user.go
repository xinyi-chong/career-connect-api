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

func (c *ControllerV1) GetScheduleByIDUser(ctx context.Context, req *v1.GetScheduleByIDUserReq) (res *v1.GetScheduleByIDUserRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.UserID == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
		return
	}
	
	r := g.RequestFromCtx(ctx)
	scheduleID := r.GetRouter("schedule_id").String()

	schedule, err := service.Application().GetScheduleByIDUser(ctx, scheduleID, *tokenData.UserID)
	
	if err != nil {
		return
	}

	res = &v1.GetScheduleByIDUserRes{
		Schedule: schedule,
	}
	
	return
}
