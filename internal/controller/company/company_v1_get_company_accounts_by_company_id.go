package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetCompanyAccountsByCompanyID(ctx context.Context, req *v1.GetCompanyAccountsByCompanyIDReq) (res *v1.GetCompanyAccountsByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	companyAccounts, err := service.Company().GetCompanyAccountsByCompanyID(ctx, companyID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyAccountsByCompanyIDRes{
		CompanyAccounts: companyAccounts,
	}

	return
}
