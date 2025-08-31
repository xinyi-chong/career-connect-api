package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateEducationByID(ctx context.Context, req *v1.PatchUpdateEducationByIDReq) (res *v1.PatchUpdateEducationByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	educationID := r.GetRouter("education_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().PatchUpdateEducationByID(ctx, req, educationID, *tokenData.UserID)

	return
}
