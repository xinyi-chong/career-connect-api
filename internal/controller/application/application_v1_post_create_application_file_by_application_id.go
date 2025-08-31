package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateApplicationFilesByApplicationID(ctx context.Context, req *v1.PostCreateApplicationFilesByApplicationIDReq) (res *v1.PostCreateApplicationFilesByApplicationIDRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	noOfSuccessResumes, noOfSuccessOtherFiles := service.Application().PostCreateApplicationFilesByApplicationID(ctx, req, applicationID, tokenData.AccountID)
	
	res = &v1.PostCreateApplicationFilesByApplicationIDRes{
		SuccessUploadedResumes: noOfSuccessResumes,
		SuccessUploadedOtherFiles: noOfSuccessOtherFiles,
	}

	return
}
