package notification

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetNotificationByID(ctx context.Context, req *v1.GetNotificationByIDReq) (res *v1.GetNotificationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	notificationID := r.GetRouter("notification_id").String()
 
	notification, err := service.Notification().GetNotificationByID(ctx, notificationID);

	res = &v1.GetNotificationByIDRes{
		Notification: notification,
	}

	return
}
