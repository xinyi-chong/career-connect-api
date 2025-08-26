package user

import (
	"context"
	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

func (s *sUser) GetUserSubscriptionByUserID(ctx context.Context, userID string) (*entity.UserSubscription, error) {
	var userSubscription *entity.UserSubscription
	err := dao.Usersubscription.Ctx(ctx).With(
		entity.UserPlan{},
	).Where(do.UserSubscription{
		UserId: userID,
	}).Scan(&userSubscription)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get User Subscription By User ID", userID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get User Subscription By User ID: " + err.Error())
	}

	return userSubscription, err
}

func (s *sUser) GetUserSubscriptionByID(ctx context.Context, userSubscriptionID string) (*entity.UserSubscription, error) {
	var userSubscription *entity.UserSubscription
	err := dao.Usersubscription.Ctx(ctx).With(
		entity.UserPlan{},
	).Where(do.UserSubscription{
		Id: userSubscriptionID,
	}).Scan(&userSubscription)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get User Subscription By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get User Subscription By ID: " + err.Error())
	}

	return userSubscription, err
}

func (s *sUser) PostCreateUserSubscription(ctx context.Context, req *v1.PostCreateUserSubscriptionReq, userID string) (*string, error) {
	userSubscriptionID := uuid.New().String()
	_, err := dao.Usersubscription.Ctx(ctx).Data(do.UserSubscription{
		Id: userSubscriptionID,
		UserId: userID,
		UserPlanId: req.UserPlanID,
		Status: consts.ACTIVE,
		Expiry: req.Expiry,
	}).Insert()
	
	if err != nil {
		g.Log().Error(ctx, "Failed to Create User Subscription: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create User Subscription: " + err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)
	
	g.Log().Info(ctx, consts.SUCCESS_CREATE, "User Subscription: ", userSubscriptionID)
	return &userSubscriptionID, err
}

func (s *sUser) PatchUpdateUserSubscriptionByUserID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionMeReq, userID string) error {
	_, err := dao.Usersubscription.Ctx(ctx).Data(do.UserSubscription{
		UserPlanId: req.UserPlanID,
		Status: req.Status,
		Expiry: req.Expiry,
	}).Where(do.UserSubscription{
		UserId: userID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update User Subscription By User ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update User Subscription By User ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "User Subscription By User ID: ", userID)
	return nil
}

func (s *sUser) PatchUpdateUserSubscriptionByID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionByIDReq, userSubscriptionID string) error {
	_, err := dao.Usersubscription.Ctx(ctx).Data(do.UserSubscription{
		UserPlanId: req.UserPlanID,
		Status: req.Status,
		Expiry: req.Expiry,
	}).Where(do.UserSubscription{
		Id: userSubscriptionID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update User Subscription By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update User Subscription By ID: " + err.Error())
		return err
	}

	// Remove Cache
	userSubscription, err := s.GetUserSubscriptionByID(ctx, userSubscriptionID)
	if err != nil {
		return err
	}
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userSubscription.UserId)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "User Subscription By ID: ", userSubscriptionID)
	return nil
}

func (s *sUser) DeleteUserSubscriptionByUserID(ctx context.Context, userID string) (err error) {
	_, err = dao.Usersubscription.Ctx(ctx).Where(do.UserSubscription{
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete User Subscription By User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete User Subscription By User ID: " + err.Error())
		return err
	}
	
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "User Subscription Me For User: ", userID)
	return nil
}

func (s *sUser) DeleteUserSubscriptionByID(ctx context.Context, userSubscriptionID string) (err error) {
	_, err = dao.Usersubscription.Ctx(ctx).Where(do.UserSubscription{
		Id: userSubscriptionID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete User Subscription By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete User Subscription By ID: " + err.Error())
		return err
	}

	// Remove Cache
	userSubscription, err := s.GetUserSubscriptionByID(ctx, userSubscriptionID)
	if err != nil {
		return err
	}
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userSubscription.UserId)
	
	g.Log().Info(ctx, consts.SUCCESS_DELETE, "User Subscription By ID: ", userSubscriptionID)
	return nil
}