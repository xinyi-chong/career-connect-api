package notification

import (
	"context"

	v1 "gf_demo/api/notification/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
)

type sNotification struct{}

func init() {
	service.RegisterNotification(New())
}

func New() *sNotification {
	return &sNotification{}
}

func (s *sNotification) GetNotificationByID(ctx context.Context, notificationID string) (*entity.Notification, error) {
	cacheNotif := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_NOTIFICATION_ID + notificationID, &entity.Notification{})
	if cacheNotif != nil {
		return cacheNotif.(*entity.Notification), nil
	}

	var notification *entity.Notification
	err := dao.Notification.Ctx(ctx).Where(do.Notification{
		Id: notificationID,
	}).Scan(&notification)

	if err != nil {
		g.Log().Error("Failed to Get Notification By ID:", notificationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Notification By ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_NOTIFICATION_ID + notificationID, notification)

	return notification, nil
}

func (s *sNotification) GetNotificationByAccountID(ctx context.Context, accountID string) ([]*entity.Notification, error) {
	cacheNotif := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + accountID, &[]*entity.Notification{})
	if cacheNotif != nil {
		return *cacheNotif.(*[]*entity.Notification), nil
	}
	
	var notifications []*entity.Notification
	err := dao.Notification.Ctx(ctx).Where(do.Notification{
		RecipientId: accountID,
	}).Scan(&notifications)

	if err != nil {
		g.Log().Error("Failed to Get Notifications By Account ID:", accountID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Notifications By Account ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + accountID, notifications)

	return notifications, nil
}

func (s *sNotification) GetNotificationByCompanyID(ctx context.Context, companyID string) ([]*entity.Notification, error) {
	company, err := service.Company().GetCompanyByAccountID(ctx, companyID)
	if err != nil {
		return nil, err
	}

	notifications, err := s.GetNotificationByAccountID(ctx, company.AccountId)

	return notifications, err
}

func (s *sNotification) GetNotificationByUserID(ctx context.Context, userID string) ([]*entity.Notification, error) {
	user, err := service.User().GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	notifications, err := s.GetNotificationByAccountID(ctx, user.AccountId)

	return notifications, err
}

func (s *sNotification) PostCreateNotification(ctx context.Context, req *v1.PostCreateNotificationReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Notification.Ctx(ctx).Data(do.Notification{
		Id: id,
		RecipientId: req.RecipientID,
		Redirect: req.Redirect,
		Title: req.Title,
		Description: req.Description,
		Status: consts.SENT,
	}).Insert()

	if err != nil {
		g.Log().Error("Failed to Create Notification", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Notification: " + err.Error())
		return nil, err
	}

	// Remove Notification Cache
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + req.RecipientID)

	g.Log().Info(consts.SUCCESS_CREATE, "Notification: ", id)
	return &id, nil
}

func (s *sNotification) PatchUpdateNotificationStatusByID(ctx context.Context, req *v1.PatchUpdateNotificationStatusByIDReq, notificationID string) (error) {
	notification, err := s.GetNotificationByID(ctx, notificationID)
	if err != nil {
		return err
	}

	_, err = dao.Notification.Ctx(ctx).Data(do.Notification{
		Status: consts.SEEN,
	}).Where(do.Notification{
		Id: notificationID,
	}).Update()
	if err != nil {
		g.Log().Error("Failed to Update Notification Status By ID", notificationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Notification Status By ID: " + err.Error())
		return err
	}

	// Remove Notification Cache
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_NOTIFICATION_ID + notificationID)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + notification.RecipientId)

	g.Log().Info(consts.SUCCESS_UPDATE, "Notification Status By ID", notificationID)
	return nil
}

func (s *sNotification) DeleteNotificationByID(ctx context.Context, notificationID string) (error) {
	notification, err := s.GetNotificationByID(ctx, notificationID)
	if err != nil {
		return err
	}

	_, err = dao.Notification.Ctx(ctx).Where(do.Notification{
		Id: notificationID,
	}).Delete()
	if err != nil {
		g.Log().Error("Failed to Delete Notification By ID", notificationID, err) 
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Notification By ID: " + err.Error())
		return err
	}

	// Remove Notification Cache
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_NOTIFICATION_ID + notificationID)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + notification.RecipientId)

	g.Log().Info(consts.SUCCESS_DELETE, "Notification By ID", notificationID)
	return nil
}