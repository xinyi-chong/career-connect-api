package account

import (
	"context"

	v1 "gf_demo/api/account/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdatePasswordMe(ctx context.Context, req *v1.PatchUpdatePasswordMeReq) (res *v1.PatchUpdatePasswordMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.Account().PatchUpdatePasswordMe(ctx, req, tokenData.AccountID)

	return
}
