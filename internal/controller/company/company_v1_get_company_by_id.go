package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetCompanyByID(ctx context.Context, req *v1.GetCompanyByIDReq) (res *v1.GetCompanyByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	company, err := service.Company().GetCompanyByID(ctx, companyID)
	if err != nil {
		return
	}

	res = &v1.GetCompanyByIDRes{
		Company: company,
	}

	return
}
