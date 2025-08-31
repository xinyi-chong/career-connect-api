package application

import (
	"context"

	v1 "gf_demo/api/application/v1"
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

func (s *sApplication) GetScheduleByCompanyID(ctx context.Context, companyID string) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		CompanyId: companyID,
	}).Scan(&schedules)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By Company ID", companyID, err)
	}

	return schedules, err
}

func (s *sApplication) GetScheduleByUserID(ctx context.Context, userID string) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		UserId: userID,
	}).Scan(&schedules)

	if err != nil {
		g.Log().Error("Failed to get schedule by UserID:", userID, err)
	}

	return schedules, err
}

func (s *sApplication) GetScheduleByID(ctx context.Context, scheduleID string) (*entity.Schedule, error) {
	var schedule *entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		Id: scheduleID,
	}).Scan(&schedule)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By ID", scheduleID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Schedule By ID: "+err.Error())
	}

	return schedule, err
}

func (s *sApplication) GetScheduleByIDCompany(ctx context.Context, scheduleID string, companyID string) (*entity.Schedule, error) {
	var schedule *entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		Id:        scheduleID,
		CompanyId: companyID,
	}).Scan(&schedule)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By ID Company", scheduleID, err)
	}

	return schedule, err
}

func (s *sApplication) GetScheduleByIDUser(ctx context.Context, scheduleID string, userID string) (*entity.Schedule, error) {
	var schedule *entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		Id:     scheduleID,
		UserId: userID,
	}).Scan(&schedule)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By ID User", scheduleID, err)
	}

	return schedule, err
}

func (s *sApplication) GetScheduleByApplicationID(ctx context.Context, applicationID string) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		ApplicationId: applicationID,
	}).Scan(&schedules)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By Application ID", applicationID, err)
	}

	return schedules, err
}

func (s *sApplication) GetScheduleByApplicationIDCompany(ctx context.Context, applicationID string, companyID string) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		ApplicationId: applicationID,
		CompanyId:     companyID,
	}).Scan(&schedules)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By Application ID Company", applicationID, err)
	}

	return schedules, err
}

func (s *sApplication) GetScheduleByApplicationIDUser(ctx context.Context, applicationID string, userID string) ([]*entity.Schedule, error) {
	var schedules []*entity.Schedule
	err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		ApplicationId: applicationID,
		UserId:        userID,
	}).Scan(&schedules)

	if err != nil {
		g.Log().Error("Failed to Get Schedule By Application ID User", applicationID, err)
	}

	return schedules, err
}

func (s *sApplication) PostCreateScheduleByApplicationID(ctx context.Context, req *v1.PostCreateScheduleByApplicationIDReq, applicationID string) (*string, error) {
	application, err := service.Application().GetApplicationByID(ctx, applicationID)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = dao.Schedule.Ctx(ctx).Data(do.Schedule{
		Id:            id,
		ApplicationId: applicationID,
		StartTime:     &req.StartTime,
		EndTime:       &req.EndTime,
		Title:         req.Title,
		CompanyId:     application.Job.CompanyId,
		UserId:        application.UserId,
		Location:      req.Location,
		Link:          req.Link,
		Status:        consts.ACTIVE,
	}).Insert()
	if err != nil {
		g.Log().Error("Failed to Create Schedule By Application ID", applicationID, err)
	}

	// Remove Caches
	RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)

	return &id, err
}

func (s *sApplication) PatchUpdateScheduleByID(ctx context.Context, req *v1.PatchUpdateScheduleByIDReq, scheduleID string) error {
	_, err := dao.Schedule.Ctx(ctx).Data(do.Schedule{
		StartTime: &req.StartTime,
		EndTime:   &req.EndTime,
		Title:     req.Title,
		Location:  req.Location,
		Link:      req.Link,
		Status:    req.Status,
	}).Where(do.Schedule{
		Id: scheduleID,
	}).Update()
	if err != nil {
		g.Log().Error("Failed to Update Schedule By ID", scheduleID, err)
	}

	// Remove Caches
	application, err := s.GetApplicationByScheduleID(ctx, scheduleID)
	if application != nil {
		RemoveApplicationCache(ctx, &application.Id, application.UserId, application.JobId)
	}

	return err
}

func (s *sApplication) DeleteScheduleByID(ctx context.Context, scheduleID string) error {
	_, err := dao.Schedule.Ctx(ctx).Where(do.Schedule{
		Id: scheduleID,
	}).Delete()
	if err != nil {
		g.Log().Error("Failed to Delete Schedule By ID", scheduleID, err)
	}

	// Remove Caches
	application, err := s.GetApplicationByScheduleID(ctx, scheduleID)
	if application != nil {
		RemoveApplicationCache(ctx, &application.Id, application.UserId, application.JobId)
	}

	return err
}
