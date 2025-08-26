package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetCompanyAccountsByUserID(ctx context.Context, req *v1.GetCompanyAccountsByUserIDReq) (res *v1.GetCompanyAccountsByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	companyAccounts, err := service.Company().GetCompanyAccountsByUserID(ctx, userID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyAccountsByUserIDRes{
		CompanyAccounts: companyAccounts,
	}

	return
}
