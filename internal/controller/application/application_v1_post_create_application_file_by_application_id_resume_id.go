package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateApplicationFileByApplicationIDResumeID(ctx context.Context, req *v1.PostCreateApplicationFileByApplicationIDResumeIDReq) (res *v1.PostCreateApplicationFileByApplicationIDResumeIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()
	resumeID := r.GetRouter("resume_id").String()
	

	applicationFileID, err := service.Application().PostCreateApplicationFileByApplicationIDResumeID(ctx, applicationID, resumeID)
	
	res = &v1.PostCreateApplicationFileByApplicationIDResumeIDRes{
		Id: applicationFileID,
	}

	return
}
