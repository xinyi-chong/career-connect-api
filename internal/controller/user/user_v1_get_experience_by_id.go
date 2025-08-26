package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetExperienceByID(ctx context.Context, req *v1.GetExperienceByIDReq) (res *v1.GetExperienceByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	experienceID := r.GetRouter("experience_id").String()

	experience, err := service.User().GetExperienceByID(ctx, experienceID);

	res = &v1.GetExperienceByIDRes{
		Experience: experience,
	}

	return
}
