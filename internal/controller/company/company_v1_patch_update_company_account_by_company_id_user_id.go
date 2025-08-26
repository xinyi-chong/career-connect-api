package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByCompanyIDUserIDReq) (res *v1.PatchUpdateCompanyAccountByCompanyIDUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()
	userID := r.GetRouter("user_id").String()

	err = service.Company().PatchUpdateCompanyAccountByCompanyIDUserID(ctx, req, companyID, userID)

	return
}
