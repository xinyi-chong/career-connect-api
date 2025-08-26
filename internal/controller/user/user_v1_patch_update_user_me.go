package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateUserMe(ctx context.Context, req *v1.PatchUpdateUserMeReq) (res *v1.PatchUpdateUserMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	in := &v1.PatchUpdateUserByIDReq{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Nationality: req.Nationality,
	}
	
	err = service.User().PatchUpdateUserByID(ctx, in, *tokenData.UserID)
	if err != nil {
		return
	}

	service.Session().ResetSessionDataByAccountID(ctx, tokenData.AccountID)

	return
}
