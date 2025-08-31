package notification

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetNotificationByUserID(ctx context.Context, req *v1.GetNotificationByUserIDReq) (res *v1.GetNotificationByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()
 
	notifications, err := service.Notification().GetNotificationByUserID(ctx, userID);

	res = &v1.GetNotificationByUserIDRes{
		Notifications: notifications,
	}

	return
}
