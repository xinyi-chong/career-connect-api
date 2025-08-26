package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateJobByID(ctx context.Context, req *v1.PatchUpdateJobByIDReq) (res *v1.PatchUpdateJobByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobID := r.GetRouter("job_id").String()
	var updatedBy string
	var updatedByType string

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.CompanyID != nil {
		updatedBy = *tokenData.CompanyID
		updatedByType = consts.COMPANY
	} else {
		updatedBy = *tokenData.UserID
		updatedByType = consts.USER
	}

	err = service.Job().PatchUpdateJobByID(ctx, req, jobID, updatedBy, updatedByType)

	return
}
