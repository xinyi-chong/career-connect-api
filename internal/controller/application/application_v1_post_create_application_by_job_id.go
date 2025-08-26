package application

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateApplicationByJobID(ctx context.Context, req *v1.PostCreateApplicationByJobIDReq) (res *v1.PostCreateApplicationByJobIDRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()

	application, err := service.Application().GetApplicationByJobIDUserID(ctx, jobID, *tokenData.UserID)
	if err != nil {
		return
	} else if application != nil {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "User has already submitted an application for this job.")
		return
	}

	applicationID, err := service.Application().PostCreateApplicationByJobID(ctx, req, jobID, *tokenData.UserID)
	if err != nil {
		return
	}

	filesInput := &v1.PostCreateApplicationFilesByApplicationIDReq{
		Resumes: req.Resumes,
		OtherFiles: req.Resumes,
	}
	
	successUploadedResumes, successUploadedOtherFiles := service.Application().PostCreateApplicationFilesByApplicationID(ctx, filesInput, *applicationID, tokenData.AccountID)

	res = &v1.PostCreateApplicationByJobIDRes{
		Id: applicationID,
		SuccessUploadedResumes: successUploadedResumes,
		SuccessUploadedOtherFiles: successUploadedOtherFiles,
	}

	return
}
