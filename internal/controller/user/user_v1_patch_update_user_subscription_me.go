package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateUserSubscriptionMe(ctx context.Context, req *v1.PatchUpdateUserSubscriptionMeReq) (res *v1.PatchUpdateUserSubscriptionMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().PatchUpdateUserSubscriptionByUserID(ctx, req, *tokenData.UserID)

	return
}
