package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateResumeByID(ctx context.Context, req *v1.PatchUpdateResumeByIDReq) (res *v1.PatchUpdateResumeByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	resumeID := r.GetRouter("resume_id").String()

	err = service.User().PatchUpdateResumeByID(ctx, req, resumeID)

	return
}
