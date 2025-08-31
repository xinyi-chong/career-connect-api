package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetEducationByID(ctx context.Context, req *v1.GetEducationByIDReq) (res *v1.GetEducationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	educationID := r.GetRouter("education_id").String()

	education, err := service.User().GetEducationByID(ctx, educationID);

	res = &v1.GetEducationByIDRes{
		Education: education,
	}

	return
}
