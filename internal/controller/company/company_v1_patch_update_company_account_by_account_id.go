package company

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateCompanyAccountByAccountID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByAccountIDReq) (res *v1.PatchUpdateCompanyAccountByAccountIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	accountID := r.GetRouter("account_id").String()
	
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	account, err := service.Account().GetAccountByID(ctx, accountID)
	if err != nil {
		return
	}

	in := &v1.PatchUpdateCompanyAccountByCompanyIDUserIDReq{
		RoleID: req.RoleID,
	}
	
	err = service.Company().PatchUpdateCompanyAccountByCompanyIDUserID(ctx, in, *tokenData.CompanyID, account.User.Id)

	return
}
