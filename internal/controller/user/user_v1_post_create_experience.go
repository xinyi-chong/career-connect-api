package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateExperience(ctx context.Context, req *v1.PostCreateExperienceReq) (res *v1.PostCreateExperienceRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	id, err := service.User().PostCreateExperience(ctx, req, *tokenData.UserID)
	
	res = &v1.PostCreateExperienceRes{
		Id: *id,
	}

	return
}
