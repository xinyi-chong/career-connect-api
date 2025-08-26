package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateCompanyPlanByID(ctx context.Context, req *v1.PatchUpdateCompanyPlanByIDReq) (res *v1.PatchUpdateCompanyPlanByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyPlanID := r.GetRouter("company_plan_id").String()

	err = service.Company().PatchUpdateCompanyPlanByID(ctx, req, companyPlanID)
	
	return
}
