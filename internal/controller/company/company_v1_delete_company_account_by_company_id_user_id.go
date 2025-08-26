package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.DeleteCompanyAccountByCompanyIDUserIDReq) (res *v1.DeleteCompanyAccountByCompanyIDUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	companyID := r.GetRouter("company_id").String()
	userID := r.GetRouter("user_id").String()

	err = service.Company().DeleteCompanyAccountByCompanyIDUserID(ctx, companyID, userID)
	
	return
}
