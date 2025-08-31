package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateResume(ctx context.Context, req *v1.PostCreateResumeReq) (res *v1.PostCreateResumeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	id, err := service.User().PostCreateResume(ctx, req, *tokenData.UserID, tokenData.AccountID)
	
	res = &v1.PostCreateResumeRes{
		Id: *id,
	}

	return
}
