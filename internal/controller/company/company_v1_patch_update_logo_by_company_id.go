package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateLogoByCompanyID(ctx context.Context, req *v1.PatchUpdateLogoByCompanyIDReq) (res *v1.PatchUpdateLogoByCompanyIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	err = service.Company().PatchUpdateLogoByCompanyID(ctx, req, companyID)

	return
}
