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

func (s *sUser) GetUserPlanByID(ctx context.Context, userPlanID string) (*entity.UserPlan, error) {
	var userPlan *entity.UserPlan
	err := dao.Userplan.Ctx(ctx).Where(do.UserPlan{
		Id: userPlanID,
	}).Scan(&userPlan)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get User Plan By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get User Plan By ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_USER_PLAN_ID + userPlanID, userPlan)

	return userPlan, err
}

func (s *sUser) PostCreateUserPlan(ctx context.Context, req *v1.PostCreateUserPlanReq) (*string, error) {
	userPlanID := uuid.New().String()
	_, err := dao.Userplan.Ctx(ctx).Data(do.UserPlan{
		Id: userPlanID,
		Name: req.Name,
	}).Insert()
	
	if err != nil {
		g.Log().Error(ctx, "Failed to Create User Plan: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create User Plan: " + err.Error())
		return nil, err
	}
	
	g.Log().Info(ctx, consts.SUCCESS_CREATE, "User Plan: ", userPlanID)
	return &userPlanID, err
}

func (s *sUser) PatchUpdateUserPlanByID(ctx context.Context, req *v1.PatchUpdateUserPlanByIDReq, userPlanID string) error {
	_, err := dao.Userplan.Ctx(ctx).Data(do.UserPlan{
		Name: req.Name,
	}).Where(do.UserPlan{
		Id: userPlanID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update User Plan By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update User Plan By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_PLAN_ID + userPlanID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "UserPlan: ", userPlanID)
	return nil
}

func (s *sUser) DeleteUserPlanByID(ctx context.Context, userPlanID string) (err error) {
	_, err = dao.Userplan.Ctx(ctx).Where(do.UserPlan{
		Id: userPlanID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete User Plan By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete User Plan By ID: " + err.Error())
		return err
	}
	
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_PLAN_ID + userPlanID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "UserPlan: ", userPlanID)
	return err
}