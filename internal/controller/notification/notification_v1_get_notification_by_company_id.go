package notification

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetNotificationByCompanyID(ctx context.Context, req *v1.GetNotificationByCompanyIDReq) (res *v1.GetNotificationByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()
 
	notifications, err := service.Notification().GetNotificationByCompanyID(ctx, companyID);
	
	res = &v1.GetNotificationByCompanyIDRes{
		Notifications: notifications,
	}

	return
}
