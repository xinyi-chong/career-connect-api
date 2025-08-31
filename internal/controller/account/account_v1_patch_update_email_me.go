package account

import (
	"context"

	v1 "gf_demo/api/account/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateEmailMe(ctx context.Context, req *v1.PatchUpdateEmailMeReq) (res *v1.PatchUpdateEmailMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	in := model.PatchUpdateEmailMeInput{
		Email:     req.Email,
		AccountID: tokenData.AccountID,
		CompanyID: tokenData.CompanyID,
		UserID:    tokenData.UserID,
	}

	newAccessToken, err := service.Account().PatchUpdateEmailMe(ctx, in)

	if err != nil {
		return
	}

	res = &v1.PatchUpdateEmailMeRes{
		AccessToken: newAccessToken,
	}

	return
}
