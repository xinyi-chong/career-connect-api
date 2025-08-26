// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package notification

import (
	"context"

	"gf_demo/api/notification/v1"
)

type INotificationV1 interface {
	PostCreateNotification(ctx context.Context, req *v1.PostCreateNotificationReq) (res *v1.PostCreateNotificationRes, err error)
	GetNotificationByID(ctx context.Context, req *v1.GetNotificationByIDReq) (res *v1.GetNotificationByIDRes, err error)
	GetNotificationByAccountID(ctx context.Context, req *v1.GetNotificationByAccountIDReq) (res *v1.GetNotificationByAccountIDRes, err error)
	GetNotificationByCompanyID(ctx context.Context, req *v1.GetNotificationByCompanyIDReq) (res *v1.GetNotificationByCompanyIDRes, err error)
	GetNotificationByUserID(ctx context.Context, req *v1.GetNotificationByUserIDReq) (res *v1.GetNotificationByUserIDRes, err error)
	PatchUpdateNotificationStatusByID(ctx context.Context, req *v1.PatchUpdateNotificationStatusByIDReq) (res *v1.PatchUpdateNotificationStatusByIDRes, err error)
	DeleteNotificationByID(ctx context.Context, req *v1.DeleteNotificationByIDReq) (res *v1.DeleteNotificationByIDRes, err error)
}
