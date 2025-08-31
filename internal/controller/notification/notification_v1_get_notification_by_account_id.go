package notification

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetNotificationByAccountID(ctx context.Context, req *v1.GetNotificationByAccountIDReq) (res *v1.GetNotificationByAccountIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	accountID := r.GetRouter("account_id").String()
 
	notifications, err := service.Notification().GetNotificationByAccountID(ctx, accountID);

	res = &v1.GetNotificationByAccountIDRes{
		Notifications: notifications,
	}

	return
}
