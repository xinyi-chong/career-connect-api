package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateEducation(ctx context.Context, req *v1.PostCreateEducationReq) (res *v1.PostCreateEducationRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	id, err := service.User().PostCreateEducation(ctx, req, *tokenData.UserID)

	res = &v1.PostCreateEducationRes{
		Id: *id,
	}

	return
}
