package application

import (
	"context"
	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

type sApplication struct{}

func init() {
	service.RegisterApplication(New())
}

func New() *sApplication {
	return &sApplication{}
}

func (s *sApplication) GetApplicationByID(ctx context.Context, applicationID string) (*entity.Application, error) {
	var application *entity.Application
	err := dao.Application.Ctx(ctx).With(
		entity.ApplicationFile{},
		entity.ApplicationFile{}.Files,
		entity.Activity{},
		entity.Schedule{},
	).Where(do.Application{
		Id: applicationID,
	}).Scan(&application)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Application By ID: ", applicationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Application By ID: "+err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_APPLICATION_ID+applicationID, application)

	return application, nil
}

func (s *sApplication) GetApplicationsByUserID(ctx context.Context, userID string) ([]*entity.Application, error) {
	var applications []*entity.Application
	err := dao.Application.Ctx(ctx).With(
		entity.Job{},
		entity.Media{},
		entity.Activity{},
		entity.Schedule{},
	).Where(do.Application{
		UserId: userID,
	}).Scan(&applications)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Applications By User ID: ", userID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Applications By User ID: "+err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_APPLICATIONS_BY_USER_ID+userID, applications)

	return applications, nil
}

func (s *sApplication) GetApplicationsByJobID(ctx context.Context, jobID string) ([]*entity.Application, error) {
	var applications []*entity.Application
	err := dao.Application.Ctx(ctx).With(
		entity.Activity{},
	).Where(do.Application{
		JobId: jobID,
	}).Scan(&applications)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Application By Job ID: ", jobID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Applications By Job ID: "+err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_APPLICATIONS_BY_JOB_ID+jobID, applications)

	return applications, nil
}

func (s *sApplication) GetApplicationByJobIDUserID(ctx context.Context, jobID string, userID string) (*entity.Application, error) {
	var application *entity.Application
	err := dao.Application.Ctx(ctx).With(
		entity.Activity{},
	).Where(do.Application{
		JobId:  jobID,
		UserId: userID,
	}).Scan(&application)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Application By Job ID & User ID: ", jobID, userID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Application By Job ID & User ID: "+err.Error())
	}

	return application, err
}

func (s *sApplication) GetApplicationsByActivityID(ctx context.Context, activityID string) ([]*entity.Application, error) {
	var application []*entity.Application
	err := dao.Application.Ctx(ctx).Where(do.Application{
		ActivityId: activityID,
	}).Scan(&application)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Applications By Activity ID: ", activityID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Applications By Activity ID: "+err.Error())
	}

	return application, err
}

func (s *sApplication) GetApplicationByScheduleID(ctx context.Context, scheduleID string) (*entity.Application, error) {
	schedule, err := s.GetScheduleByID(ctx, scheduleID)
	if err != nil {
		return nil, err
	}

	application, err := s.GetApplicationByID(ctx, schedule.ApplicationId)
	return application, err
}

func (s *sApplication) PostCreateApplicationByJobID(ctx context.Context, req *v1.PostCreateApplicationByJobIDReq, jobID string, userID string) (*string, error) {
	id := uuid.New().String()

	_, err := dao.Application.Ctx(ctx).Data(do.Application{
		Id:         id,
		JobId:      jobID,
		UserId:     userID,
		Answer:     req.Answer,
		ApplyAt:    gtime.New(),
		ActivityId: req.ActivityID,
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "Failed to Create Application By Job ID: ", jobID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Application By Job ID: "+err.Error())
		return nil, err
	}

	// Remove Caches
	RemoveApplicationCache(ctx, nil, userID, jobID)

	return &id, nil
}

func (s *sApplication) PatchUpdateApplicationByID(ctx context.Context, req *v1.PatchUpdateApplicationByIDReq, applicationID string) error {
	_, err := dao.Application.Ctx(ctx).Data(do.Application{
		Answer:     req.Answer,
		ActivityId: req.ActivityID,
	}).Where(do.Application{
		Id: applicationID,
	}).Update()
	if err != nil {
		g.Log().Error(ctx, "Failed to Update Application By ID: ", applicationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Application By ID: "+err.Error())
		return err
	}

	// Remove Caches
	application, _ := s.GetApplicationByID(ctx, applicationID)
	if application != nil {
		RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Application By ID", applicationID)
	return nil
}

func (s *sApplication) DeleteApplicationByID(ctx context.Context, applicationID string) error {
	application, err := s.GetApplicationByID(ctx, applicationID)
	if err != nil {
		return err
	}

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Delete Media (Resume)
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			for _, applicationFile := range application.ApplicationFiles {
				err = service.Media().DeleteMediaByID(ctx, applicationFile.Id)
				if err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}

		if err := dao.Application.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.Application.Ctx(ctx).Where(do.Application{
				Id: applicationID,
			}).Delete()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Delete Application By ID: ", applicationID, err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Application By ID: "+err.Error())
			return err
		}

		// Remove Caches
		RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)

		g.Log().Info(ctx, consts.SUCCESS_DELETE, "Application By ID", applicationID)
		return nil
	})

	return nil
}

func RemoveApplicationCache(ctx context.Context, applicationID *string, userID string, jobID string) {
	if applicationID != nil {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_APPLICATION_ID+*applicationID)
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_APPLICATIONS_BY_JOB_ID+jobID)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_APPLICATIONS_BY_USER_ID+userID)
}
