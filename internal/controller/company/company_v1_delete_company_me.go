package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteCompanyMe(ctx context.Context, req *v1.DeleteCompanyMeReq) (res *v1.DeleteCompanyMeRes, err error) {
	err = service.Company().DeleteCompanyMe(ctx)

	return
}
