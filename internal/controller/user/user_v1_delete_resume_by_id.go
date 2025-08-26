package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteResumeByID(ctx context.Context, req *v1.DeleteResumeByIDReq) (res *v1.DeleteResumeByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	resumeID := r.GetRouter("resume_id").String()

	err = service.User().DeleteResumeByID(ctx, resumeID)

	return
}
