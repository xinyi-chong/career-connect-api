package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateProfilePicture(ctx context.Context, req *v1.PostCreateProfilePictureReq) (res *v1.PostCreateProfilePictureRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	id, err := service.User().PostCreateProfilePicture(ctx, req, *tokenData.UserID, tokenData.AccountID)

	res = &v1.PostCreateProfilePictureRes{
		Id: *id,
	}

	return
}
