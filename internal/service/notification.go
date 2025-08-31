// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/model/entity"
)

type (
	INotification interface {
		GetNotificationByID(ctx context.Context, notificationID string) (*entity.Notification, error)
		GetNotificationByAccountID(ctx context.Context, accountID string) ([]*entity.Notification, error)
		GetNotificationByCompanyID(ctx context.Context, companyID string) ([]*entity.Notification, error)
		GetNotificationByUserID(ctx context.Context, userID string) ([]*entity.Notification, error)
		PostCreateNotification(ctx context.Context, req *v1.PostCreateNotificationReq) (*string, error)
		PatchUpdateNotificationStatusByID(ctx context.Context, req *v1.PatchUpdateNotificationStatusByIDReq, notificationID string) error
		DeleteNotificationByID(ctx context.Context, notificationID string) error
	}
)

var (
	localNotification INotification
)

func Notification() INotification {
	if localNotification == nil {
		panic("implement not found for interface INotification, forgot register?")
	}
	return localNotification
}

func RegisterNotification(i INotification) {
	localNotification = i
}
