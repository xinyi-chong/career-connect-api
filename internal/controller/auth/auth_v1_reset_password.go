package auth

import (
	"context"

	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error) {
	token := g.RequestFromCtx(ctx).GetQuery("token").String()
	
	claim, err := service.Token().ParseValidateAccountToken(token, consts.RESET_PASSWORD_TOKEN)
	if err != nil {
		return
	}
	
	err = service.Account().UpdatePasswordByAccountID(ctx, req.Password, claim.AccountID)

	return
}
