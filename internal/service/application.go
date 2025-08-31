// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IApplication interface {
		GetActivityByID(ctx context.Context, activityID string) (*entity.Activity, error)
		PostCreateActivity(ctx context.Context, req *v1.PostCreateActivityReq) (*string, error)
		PostCreateActivityByApplicationID(ctx context.Context, req *v1.PostCreateActivityByApplicationIDReq, applicationID string) (activityID *string, err error)
		PatchUpdateActivityByID(ctx context.Context, req *v1.PatchUpdateActivityByIDReq, activityID string) error
		DeleteActivityByID(ctx context.Context, activityID string) error
		GetApplicationByID(ctx context.Context, applicationID string) (*entity.Application, error)
		GetApplicationsByUserID(ctx context.Context, userID string) ([]*entity.Application, error)
		GetApplicationsByJobID(ctx context.Context, jobID string) ([]*entity.Application, error)
		GetApplicationByJobIDUserID(ctx context.Context, jobID string, userID string) (*entity.Application, error)
		GetApplicationsByActivityID(ctx context.Context, activityID string) ([]*entity.Application, error)
		GetApplicationByScheduleID(ctx context.Context, scheduleID string) (*entity.Application, error)
		PostCreateApplicationByJobID(ctx context.Context, req *v1.PostCreateApplicationByJobIDReq, jobID string, userID string) (*string, error)
		PatchUpdateApplicationByID(ctx context.Context, req *v1.PatchUpdateApplicationByIDReq, applicationID string) error
		DeleteApplicationByID(ctx context.Context, applicationID string) error
		GetApplicationChatByApplicationID(ctx context.Context, applicationID string) ([]*entity.ApplicationChatMessage, error)
		PostCreateApplicationChatByApplicationID(ctx context.Context, req *v1.PostCreateApplicationChatByApplicationIDReq, applicationID string) (*string, error)
		DeleteApplicationChatByApplicationID(ctx context.Context, applicationID string) error
		GetApplicationFileByApplicationIDFileID(ctx context.Context, applicationID string, fileID string) (*entity.ApplicationFile, error)
		GetApplicationFilesByApplicationID(ctx context.Context, applicationID string) ([]*entity.ApplicationFile, error)
		PostCreateApplicationFileByApplicationIDResumeID(ctx context.Context, applicationID string, resumeID string) (*string, error)
		PostCreateApplicationFilesByApplicationID(ctx context.Context, req *v1.PostCreateApplicationFilesByApplicationIDReq, applicationID string, accountID string) (int, int)
		UploadMultileApplicationFiles(ctx context.Context, files []*ghttp.UploadFile, fileType string, accountID string, applicationID string) int
		DeleteApplicationFileByApplicationIDFileID(ctx context.Context, applicationID string, fileID string) error
		GetScheduleByCompanyID(ctx context.Context, companyID string) ([]*entity.Schedule, error)
		GetScheduleByUserID(ctx context.Context, userID string) ([]*entity.Schedule, error)
		GetScheduleByID(ctx context.Context, scheduleID string) (*entity.Schedule, error)
		GetScheduleByIDCompany(ctx context.Context, scheduleID string, companyID string) (*entity.Schedule, error)
		GetScheduleByIDUser(ctx context.Context, scheduleID string, userID string) (*entity.Schedule, error)
		GetScheduleByApplicationID(ctx context.Context, applicationID string) ([]*entity.Schedule, error)
		GetScheduleByApplicationIDCompany(ctx context.Context, applicationID string, companyID string) ([]*entity.Schedule, error)
		GetScheduleByApplicationIDUser(ctx context.Context, applicationID string, userID string) ([]*entity.Schedule, error)
		PostCreateScheduleByApplicationID(ctx context.Context, req *v1.PostCreateScheduleByApplicationIDReq, applicationID string) (*string, error)
		PatchUpdateScheduleByID(ctx context.Context, req *v1.PatchUpdateScheduleByIDReq, scheduleID string) error
		DeleteScheduleByID(ctx context.Context, scheduleID string) error
	}
)

var (
	localApplication IApplication
)

func Application() IApplication {
	if localApplication == nil {
		panic("implement not found for interface IApplication, forgot register?")
	}
	return localApplication
}

func RegisterApplication(i IApplication) {
	localApplication = i
}
