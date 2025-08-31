package application

import (
	"context"
	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

func (s *sApplication) GetActivityByID(ctx context.Context, activityID string) (*entity.Activity, error) {
	var activity *entity.Activity
	err := dao.Activity.Ctx(ctx).Where(do.Activity{
		Id: activityID,
	}).Scan(&activity)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Activity By ID: ", activityID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Activity By ID: "+err.Error())
		return nil, err
	}

	return activity, nil
}

func (s *sApplication) PostCreateActivity(ctx context.Context, req *v1.PostCreateActivityReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Activity.Ctx(ctx).Data(do.Activity{
		Id:   id,
		Name: req.Name,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Activity: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Activity: "+err.Error())
		return nil, err
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Activity: ", id, err)
	return &id, nil
}

func (s *sApplication) PostCreateActivityByApplicationID(ctx context.Context, req *v1.PostCreateActivityByApplicationIDReq, applicationID string) (activityID *string, err error) {
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Activity
		if err := dao.Activity.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			activityID, err = s.PostCreateActivity(ctx, &v1.PostCreateActivityReq{Name: req.Name})
			return err
		}); err != nil {
			return err
		}

		// Update Application With New Activity ID
		if err := dao.Application.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := dao.Application.Ctx(ctx).Data(do.Application{
				ActivityId: activityID,
			}).Where(do.Application{
				Id: applicationID,
			}).Update()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Update Application's Activity ID By Application ID: ", applicationID, err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Application's Activity ID By Application ID: "+err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Remove Caches
	application, err := s.GetApplicationByID(ctx, applicationID)
	if application != nil {
		RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Activity By Application ID: ", applicationID, err)
	return activityID, nil
}

func (s *sApplication) PatchUpdateActivityByID(ctx context.Context, req *v1.PatchUpdateActivityByIDReq, activityID string) error {
	_, err := dao.Activity.Ctx(ctx).Data(do.Activity{
		Name: req.Name,
	}).Where(do.Activity{
		Id: activityID,
	}).Update()

	if err != nil {
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Activity By ID: "+err.Error())
		g.Log().Error(ctx, "Failed to Update Activity By ID: ", activityID, err)
		return err
	}

	// Remove Caches
	applications, _ := s.GetApplicationsByActivityID(ctx, activityID)
	for _, application := range applications {
		RemoveApplicationCache(ctx, &application.Id, application.UserId, application.JobId)
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Activity By ID: ", activityID, err)
	return nil
}

func (s *sApplication) DeleteActivityByID(ctx context.Context, activityID string) error {
	_, err := dao.Activity.Ctx(ctx).Where(do.Activity{
		Id: activityID,
	}).Delete()

	if err != nil {
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Activity By ID: "+err.Error())
		g.Log().Error(ctx, "Failed to Delete Activity By ID: ", activityID, err)
		return err
	}

	// Remove Caches
	applications, _ := s.GetApplicationsByActivityID(ctx, activityID)
	for _, application := range applications {
		RemoveApplicationCache(ctx, &application.Id, application.UserId, application.JobId)
	}

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Activity By ID", activityID)
	return nil
}
