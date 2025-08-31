package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetResumeByID(ctx context.Context, req *v1.GetResumeByIDReq) (res *v1.GetResumeByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	resumeID := r.GetRouter("resume_id").String()

	resume, err := service.User().GetResumeByID(ctx, resumeID);
	
	res = &v1.GetResumeByIDRes{
		Resume: resume,
	}

	return
}
