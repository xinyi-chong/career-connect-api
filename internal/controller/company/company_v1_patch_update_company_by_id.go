package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateCompanyByID(ctx context.Context, req *v1.PatchUpdateCompanyByIDReq) (res *v1.PatchUpdateCompanyByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()

	err = service.Company().PatchUpdateCompanyByID(ctx, req, companyID)

	return
}
