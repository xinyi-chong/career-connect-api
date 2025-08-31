package notification

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteNotificationByID(ctx context.Context, req *v1.DeleteNotificationByIDReq) (res *v1.DeleteNotificationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	notificationID := r.GetRouter("notification_id").String()
 
	err = service.Notification().DeleteNotificationByID(ctx, notificationID);

	return
}
