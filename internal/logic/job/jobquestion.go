package job

import (
	"context"
	v1 "gf_demo/api/job/v1"
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

func (s *sJob) GetJobQuestionByID(ctx context.Context, jobQuestionID string) (*entity.JobQuestion, error) {
	var jobQuestion *entity.JobQuestion
	err := dao.Jobquestion.Ctx(ctx).Where(do.JobQuestion{
		Id: jobQuestionID,
	}).Scan(&jobQuestion)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Job Question by ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Job Question By ID: " + err.Error())
	}
	
	return jobQuestion, err
}

func (s *sJob) GetJobQuestionByJobID(ctx context.Context, jobID string) (*entity.JobQuestion, error) {
	var jobQuestion *entity.JobQuestion
	err := dao.Jobquestion.Ctx(ctx).Where(do.JobQuestion{
		JobId: jobID,
	}).Scan(&jobQuestion)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Job Question By Job ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Job Question By Job ID: " + err.Error())
	}

	return jobQuestion, err
}

func (s *sJob) PostCreateJobQuestion(ctx context.Context, req *v1.PostCreateJobQuestionReq) (*string, error) {
	id := uuid.New().String()
	
	_, err := dao.Jobquestion.Ctx(ctx).Data(do.JobQuestion{
		Id: id,
		Question: req.Question,
		JobId: req.JobID,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Job Question: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Job Question By ID: " + err.Error())
		return nil, err
	}

	job, _ := s.GetJobByID(ctx, req.JobID)
	if job != nil {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS_BY_COMPANY_ID + job.CompanyId)
	}
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOB_ID + req.JobID)

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Job Question, ID: ", id)
	return &id, nil
}

func (s *sJob) PatchUpdateJobQuestionByID(ctx context.Context, req *v1.PatchUpdateJobQuestionByIDReq, id string) (error) {
	_, err := dao.Jobquestion.Ctx(ctx).Data(do.JobQuestion{
		Question: req.Question,
	}).Where(do.JobQuestion{
		Id: id,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Job Question By ID: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Job Question By ID: " + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Job Question")

	jobQuestion, _ := s.GetJobQuestionByID(ctx, id)
	if jobQuestion != nil {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOB_ID + jobQuestion.JobId)
		job, _ := s.GetJobByID(ctx, jobQuestion.JobId)
		if job != nil {
			service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS_BY_COMPANY_ID + job.CompanyId)
		}
	}
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS)

	return nil
}

func (s *sJob) DeleteJobQuestionByID(ctx context.Context, id string) error {
	jobQuestion, err := s.GetJobQuestionByID(ctx, id)
	if err != nil {
		return err
	}

	job, err := s.GetJobByID(ctx, jobQuestion.JobId)
	if err != nil {
		return err
	}

	_, err = dao.Jobquestion.Ctx(ctx).Where(do.JobQuestion{
		Id: id,
	}).Delete()
	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Job Question By ID: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Job Question By ID: " + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Job Question")

	cacheKeys := []string{
		consts.CACHE_JOB_ID + job.Id,
		consts.CACHE_JOBS,
		consts.CACHE_JOBS_BY_COMPANY_ID + job.CompanyId,
	}
	service.Cache().RemoveMulCachesWithPrefix(ctx, cacheKeys)
	
	return nil
}