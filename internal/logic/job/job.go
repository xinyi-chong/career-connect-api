package job

import (
	"context"
	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sJob struct{}

func init() {
	service.RegisterJob(New())
}

func New() *sJob {
	return &sJob{}
}

func (s *sJob) GetJobByID(ctx context.Context, jobID string) (*entity.Job, error) {
	cacheJob := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_JOB_ID + jobID, &entity.Job{})
	if cacheJob != nil {
		return cacheJob.(*entity.Job), nil
	}
	
	var job *entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Where(do.Job{
		Id: jobID,
	}).Scan(&job)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Job by ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Job by ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_JOB_ID + jobID, job)
	
	return job, err
}

func (s *sJob) GetJobsAll(ctx context.Context) ([]*entity.Job, error) {
	cacheJobs := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_JOBS, &[]*entity.Job{})
	if cacheJobs != nil {
		return *cacheJobs.(*[]*entity.Job), nil
	}

	var jobs []*entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Scan(&jobs)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get All Jobs", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Jobs All: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_JOBS, jobs)

	return jobs, nil
}

func (s *sJob) GetJobsAllByKeywords(ctx context.Context, keywords []interface{}) ([]*entity.Job, error) {
	
	conditions := make([]string, len(keywords))
	args := make([]interface{}, len(keywords))
	for i, keyword := range keywords {
		// Use FIND_IN_SET to check if the tag column contains the keyword
		conditions[i] = "FIND_IN_SET(?, tag) > 0"
		args[i] = keyword
	}

	whereClause := strings.Join(conditions, " OR ")

	var jobs []*entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Where(whereClause, args...).Scan(&jobs)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Jobs By Keywords", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Jobs By Keywords: " + err.Error())
	}
	
	return jobs, err
}

func (s *sJob) GetJobsByCompanyID(ctx context.Context, companyID string) ([]*entity.Job, error) {
	cacheJob := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_JOBS_BY_COMPANY_ID + companyID, &[]*entity.Job{})
	if cacheJob != nil {
		return *cacheJob.(*[]*entity.Job), nil
	}
	
	var jobs []*entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Where(do.Job{
		CompanyId: companyID,
	}).Scan(&jobs)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Jobs By Company ID: ", companyID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Jobs By Company ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_JOBS_BY_COMPANY_ID + companyID, jobs)

	return jobs, nil
}

func (s *sJob) GetJobsByUserID(ctx context.Context, userID string) (map[string][]*entity.Job, error) {	
	companyAccounts, err := service.Company().GetCompanyAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	userJobs := make(map[string][]*entity.Job)

	for _, companyAccount := range companyAccounts {
		var jobs []*entity.Job
		jobs, err = s.GetJobsByCompanyID(ctx, companyAccount.CompanyId)
		if err != nil {
			return nil, err
		}

		userJobs[companyAccount.CompanyId] = jobs
	}
	
	return userJobs, nil
}

func (s *sJob) GetJobsCreatedByCompanyID(ctx context.Context, companyID string) ([]*entity.Job, error) {
	var jobs []*entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Where(do.Job{
		CreatedBy: companyID,
	}).Scan(&jobs)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Jobs Created By Company ID: ", companyID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Jobs Created By Company ID: " + err.Error())
	}
	
	return jobs, err
}

func (s *sJob) GetJobsCreatedByUserID(ctx context.Context, userID string) ([]*entity.Job, error) {
	var jobs []*entity.Job
	err := dao.Job.Ctx(ctx).With(
		entity.JobQuestion{},
	).Where(do.Job{
		CreatedBy: userID,
	}).Scan(&jobs)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Jobs Created By User ID: ", userID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Jobs Created By User ID: " + err.Error())
	}
	
	return jobs, err
}

func (s *sJob) PostCreateJob(ctx context.Context, req model.PostCreateJobInput) (id *string, err error) {
	jobID := uuid.New().String()

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		var jobQuestionID *string

		// Insert Job
		if err = dao.Job.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := dao.Job.Ctx(ctx).Data(do.Job{
				Id: jobID,
				Title: req.Title,
				CompanyId: req.CompanyID,
				Tag: req.Tag,
				Description: req.Description,
				Level: req.Level,
				Salary: req.Salary,
				PostedAt: &req.PostedAt,
				Location: req.Location,
				IsRemote: req.IsRemote,
				IsHybrid: req.IsHybrid,
				Expiry: &req.Expiry,
				Status: consts.ACTIVE,
				CreatedBy: req.CreatedBy,
				CreatedByType: req.CreatedByType,
				UpdatedBy: req.UpdatedBy,
				UpdatedByType: req.UpdatedByType,
			}).Insert()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Create Job: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Job: " + err.Error())
			return err
		}

		if req.JobQuestion != nil {
			// Insert Job Question
			if err := dao.Jobquestion.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				jobQuestion := &v1.PostCreateJobQuestionReq{
					Question: *req.JobQuestion,
					JobID: jobID,
				}
				jobQuestionID, err = s.PostCreateJobQuestion(ctx, jobQuestion)
				return err
			}); err != nil {
				return err
			}
		}

		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS)
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_JOBS_BY_COMPANY_ID + req.CompanyID)

		g.Log().Info(ctx, consts.SUCCESS_CREATE, "Job ID: ", jobID, " Job Question ID: ", jobQuestionID)
		return nil
	})

	return &jobID, err
}

func (s *sJob) PatchUpdateJobByID(ctx context.Context, req *v1.PatchUpdateJobByIDReq, id string, updatedBy string, updatedByType string) (error) {
	job, err := s.GetJobByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = dao.Job.Ctx(ctx).Data(do.Job{
		Title: req.Title,
		Tag: req.Tag,
		Description: req.Description,
		Level: req.Level,
		Salary: req.Salary,
		Location: req.Location,
		IsRemote: req.IsRemote,
		IsHybrid: req.IsHybrid,
		Expiry: &req.Expiry,
		Status: req.Status,
		UpdatedBy: updatedBy,
		UpdatedByType: updatedByType,
	}).Where(do.Job{
		Id: id,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Job By ID: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Job By ID: " + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Job")

	//Remove Caches
	cacheKeys := []string{
		consts.CACHE_JOB_ID + id,
		consts.CACHE_JOBS,
		consts.CACHE_JOBS_BY_COMPANY_ID + job.CompanyId,
	}
	service.Cache().RemoveMulCachesWithPrefix(ctx, cacheKeys)

	return err
}

func (s *sJob) DeleteJobByID(ctx context.Context, id string) (error) {
	job, err := s.GetJobByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = dao.Job.Ctx(ctx).Where(do.Job{
		Id: id,
	}).Delete()
	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Job By ID: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Job By ID: " + err.Error())
		return err
	}
	
	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Job")

	//Remove Caches
	cacheKeys := []string{
		consts.CACHE_JOB_ID + id,
		consts.CACHE_JOBS,
		consts.CACHE_JOBS_BY_COMPANY_ID + job.CompanyId,
		consts.CACHE_APPLICATIONS_BY_JOB_ID + id,
	}
	service.Cache().RemoveMulCachesWithPrefix(ctx, cacheKeys)

	return nil
}