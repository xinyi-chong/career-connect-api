package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Validate(ctx context.Context, req *v1.ValidateReq) (res *v1.ValidateRes, err error) {
	token := g.RequestFromCtx(ctx).GetQuery("token").String()
	
	claim, err := service.Token().ParseValidateAccountToken(token, consts.VALIDATE_ACCOUNT_TOKEN)
	if err != nil {
		return
	}

	err = service.Auth().Validate(ctx, claim.AccountID)

	return
}
