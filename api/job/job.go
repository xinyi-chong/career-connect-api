// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package job

import (
	"context"

	"gf_demo/api/job/v1"
)

type IJobV1 interface {
	PostCreateJob(ctx context.Context, req *v1.PostCreateJobReq) (res *v1.PostCreateJobRes, err error)
	GetJobsAllMe(ctx context.Context, req *v1.GetJobsAllMeReq) (res *v1.GetJobsAllMeRes, err error)
	GetJobsAllByKeywords(ctx context.Context, req *v1.GetJobsAllByKeywordsReq) (res *v1.GetJobsAllByKeywordsRes, err error)
	GetJobsByCompanyID(ctx context.Context, req *v1.GetJobsByCompanyIDReq) (res *v1.GetJobsByCompanyIDRes, err error)
	GetJobsMeCompany(ctx context.Context, req *v1.GetJobsMeCompanyReq) (res *v1.GetJobsMeCompanyRes, err error)
	GetJobsMeUser(ctx context.Context, req *v1.GetJobsMeUserReq) (res *v1.GetJobsMeUserRes, err error)
	GetJobByIDCompany(ctx context.Context, req *v1.GetJobByIDCompanyReq) (res *v1.GetJobByIDCompanyRes, err error)
	GetJobByIDUser(ctx context.Context, req *v1.GetJobByIDUserReq) (res *v1.GetJobByIDUserRes, err error)
	GetJobsCreatedByCompanyID(ctx context.Context, req *v1.GetJobsCreatedByCompanyIDReq) (res *v1.GetJobsCreatedByCompanyIDRes, err error)
	GetJobsCreatedByUserID(ctx context.Context, req *v1.GetJobsCreatedByUserIDReq) (res *v1.GetJobsCreatedByUserIDRes, err error)
	PatchUpdateJobByID(ctx context.Context, req *v1.PatchUpdateJobByIDReq) (res *v1.PatchUpdateJobByIDRes, err error)
	DeleteJobByID(ctx context.Context, req *v1.DeleteJobByIDReq) (res *v1.DeleteJobByIDRes, err error)
	PostCreateJobQuestion(ctx context.Context, req *v1.PostCreateJobQuestionReq) (res *v1.PostCreateJobQuestionRes, err error)
	GetJobQuestionByID(ctx context.Context, req *v1.GetJobQuestionByIDReq) (res *v1.GetJobQuestionByIDRes, err error)
	PatchUpdateJobQuestionByID(ctx context.Context, req *v1.PatchUpdateJobQuestionByIDReq) (res *v1.PatchUpdateJobQuestionByIDRes, err error)
	DeleteJobQuestionByID(ctx context.Context, req *v1.DeleteJobQuestionByIDReq) (res *v1.DeleteJobQuestionByIDRes, err error)
}
