package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteEducationByID(ctx context.Context, req *v1.DeleteEducationByIDReq) (res *v1.DeleteEducationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	educationID := r.GetRouter("education_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().DeleteEducationByID(ctx, educationID, *tokenData.UserID)

	return
}
