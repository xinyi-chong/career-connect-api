// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package application

import (
	"context"

	"gf_demo/api/application/v1"
)

type IApplicationV1 interface {
	PostCreateApplicationByJobID(ctx context.Context, req *v1.PostCreateApplicationByJobIDReq) (res *v1.PostCreateApplicationByJobIDRes, err error)
	GetApplicationByJobID(ctx context.Context, req *v1.GetApplicationByJobIDReq) (res *v1.GetApplicationByJobIDRes, err error)
	GetApplicationsMe(ctx context.Context, req *v1.GetApplicationsMeReq) (res *v1.GetApplicationsMeRes, err error)
	GetApplicationByID(ctx context.Context, req *v1.GetApplicationByIDReq) (res *v1.GetApplicationByIDRes, err error)
	PatchUpdateApplicationByID(ctx context.Context, req *v1.PatchUpdateApplicationByIDReq) (res *v1.PatchUpdateApplicationByIDRes, err error)
	DeleteApplicationByID(ctx context.Context, req *v1.DeleteApplicationByIDReq) (res *v1.DeleteApplicationByIDRes, err error)
	PostCreateActivity(ctx context.Context, req *v1.PostCreateActivityReq) (res *v1.PostCreateActivityRes, err error)
	PostCreateActivityByApplicationID(ctx context.Context, req *v1.PostCreateActivityByApplicationIDReq) (res *v1.PostCreateActivityByApplicationIDRes, err error)
	GetActivityByID(ctx context.Context, req *v1.GetActivityByIDReq) (res *v1.GetActivityByIDRes, err error)
	PatchUpdateActivityByID(ctx context.Context, req *v1.PatchUpdateActivityByIDReq) (res *v1.PatchUpdateActivityByIDRes, err error)
	DeleteActivityByID(ctx context.Context, req *v1.DeleteActivityByIDReq) (res *v1.DeleteActivityByIDRes, err error)
	PostCreateApplicationChatByApplicationID(ctx context.Context, req *v1.PostCreateApplicationChatByApplicationIDReq) (res *v1.PostCreateApplicationChatByApplicationIDRes, err error)
	GetApplicationChatByApplicationID(ctx context.Context, req *v1.GetApplicationChatByApplicationIDReq) (res *v1.GetApplicationChatByApplicationIDRes, err error)
	DeleteApplicationChatByApplicationID(ctx context.Context, req *v1.DeleteApplicationChatByApplicationIDReq) (res *v1.DeleteApplicationChatByApplicationIDRes, err error)
	PostCreateApplicationFilesByApplicationID(ctx context.Context, req *v1.PostCreateApplicationFilesByApplicationIDReq) (res *v1.PostCreateApplicationFilesByApplicationIDRes, err error)
	PostCreateApplicationFileByApplicationIDResumeID(ctx context.Context, req *v1.PostCreateApplicationFileByApplicationIDResumeIDReq) (res *v1.PostCreateApplicationFileByApplicationIDResumeIDRes, err error)
	GetApplicationFilesByApplicationID(ctx context.Context, req *v1.GetApplicationFilesByApplicationIDReq) (res *v1.GetApplicationFilesByApplicationIDRes, err error)
	DeleteApplicationFileByApplicationIDFileID(ctx context.Context, req *v1.DeleteApplicationFileByApplicationIDFileIDReq) (res *v1.DeleteApplicationFileByApplicationIDFileIDRes, err error)
	PostCreateScheduleByApplicationID(ctx context.Context, req *v1.PostCreateScheduleByApplicationIDReq) (res *v1.PostCreateScheduleByApplicationIDRes, err error)
	GetScheduleMeCompany(ctx context.Context, req *v1.GetScheduleMeCompanyReq) (res *v1.GetScheduleMeCompanyRes, err error)
	GetScheduleMeUser(ctx context.Context, req *v1.GetScheduleMeUserReq) (res *v1.GetScheduleMeUserRes, err error)
	GetScheduleByCompanyID(ctx context.Context, req *v1.GetScheduleByCompanyIDReq) (res *v1.GetScheduleByCompanyIDRes, err error)
	GetScheduleByApplicationID(ctx context.Context, req *v1.GetScheduleByApplicationIDReq) (res *v1.GetScheduleByApplicationIDRes, err error)
	GetScheduleByApplicationIDCompany(ctx context.Context, req *v1.GetScheduleByApplicationIDCompanyReq) (res *v1.GetScheduleByApplicationIDCompanyRes, err error)
	GetScheduleByApplicationIDUser(ctx context.Context, req *v1.GetScheduleByApplicationIDUserReq) (res *v1.GetScheduleByApplicationIDUserRes, err error)
	GetScheduleByIDCompany(ctx context.Context, req *v1.GetScheduleByIDCompanyReq) (res *v1.GetScheduleByIDCompanyRes, err error)
	GetScheduleByIDUser(ctx context.Context, req *v1.GetScheduleByIDUserReq) (res *v1.GetScheduleByIDUserRes, err error)
	PatchUpdateScheduleByID(ctx context.Context, req *v1.PatchUpdateScheduleByIDReq) (res *v1.PatchUpdateScheduleByIDRes, err error)
	DeleteScheduleByID(ctx context.Context, req *v1.DeleteScheduleByIDReq) (res *v1.DeleteScheduleByIDRes, err error)
}
