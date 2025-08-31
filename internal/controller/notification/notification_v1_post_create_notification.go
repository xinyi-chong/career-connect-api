package notification

import (
	"context"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateNotification(ctx context.Context, req *v1.PostCreateNotificationReq) (res *v1.PostCreateNotificationRes, err error) {
	notificationID, err := service.Notification().PostCreateNotification(ctx, req);

	res = &v1.PostCreateNotificationRes{
		Id: notificationID,
	}

	return
}
