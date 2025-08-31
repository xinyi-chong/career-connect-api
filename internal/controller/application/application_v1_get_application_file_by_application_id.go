package application

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetApplicationFilesByApplicationID(ctx context.Context, req *v1.GetApplicationFilesByApplicationIDReq) (res *v1.GetApplicationFilesByApplicationIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	applicationID := r.GetRouter("application_id").String()

	applicationFiles, err := service.Application().GetApplicationFilesByApplicationID(ctx, applicationID)

	if err != nil {
		return
	}

	var resumes []*entity.ApplicationFile
	var otherFiles []*entity.ApplicationFile
	for _, applicationFile := range applicationFiles {
		if applicationFile.FileType == consts.RESUME {
			resumes = append(resumes, applicationFile)
		} else {
			otherFiles = append(otherFiles, applicationFile)
		}
	}

	res = &v1.GetApplicationFilesByApplicationIDRes{
		Resumes: resumes,
		OtherFiles: otherFiles,
	}

	return
}
