package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteCompanyAccountByAccountID(ctx context.Context, req *v1.DeleteCompanyAccountByAccountIDReq) (res *v1.DeleteCompanyAccountByAccountIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	accountID := r.GetRouter("account_id").String()

	account, err := service.Account().GetAccountByID(ctx, accountID)
	if err != nil {
		return
	}

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.Company().DeleteCompanyAccountByCompanyIDUserID(ctx, *tokenData.CompanyID, account.User.Id)

	return
}
