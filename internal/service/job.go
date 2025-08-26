// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
)

type (
	IJob interface {
		GetJobByID(ctx context.Context, jobID string) (*entity.Job, error)
		GetJobsAll(ctx context.Context) ([]*entity.Job, error)
		GetJobsAllByKeywords(ctx context.Context, keywords []interface{}) ([]*entity.Job, error)
		GetJobsByCompanyID(ctx context.Context, companyID string) ([]*entity.Job, error)
		GetJobsByUserID(ctx context.Context, userID string) (map[string][]*entity.Job, error)
		GetJobsCreatedByCompanyID(ctx context.Context, companyID string) ([]*entity.Job, error)
		GetJobsCreatedByUserID(ctx context.Context, userID string) ([]*entity.Job, error)
		PostCreateJob(ctx context.Context, req model.PostCreateJobInput) (id *string, err error)
		PatchUpdateJobByID(ctx context.Context, req *v1.PatchUpdateJobByIDReq, id string, updatedBy string, updatedByType string) error
		DeleteJobByID(ctx context.Context, id string) error
		GetJobQuestionByID(ctx context.Context, jobQuestionID string) (*entity.JobQuestion, error)
		GetJobQuestionByJobID(ctx context.Context, jobID string) (*entity.JobQuestion, error)
		PostCreateJobQuestion(ctx context.Context, req *v1.PostCreateJobQuestionReq) (*string, error)
		PatchUpdateJobQuestionByID(ctx context.Context, req *v1.PatchUpdateJobQuestionByIDReq, id string) error
		DeleteJobQuestionByID(ctx context.Context, id string) error
	}
)

var (
	localJob IJob
)

func Job() IJob {
	if localJob == nil {
		panic("implement not found for interface IJob, forgot register?")
	}
	return localJob
}

func RegisterJob(i IJob) {
	localJob = i
}
