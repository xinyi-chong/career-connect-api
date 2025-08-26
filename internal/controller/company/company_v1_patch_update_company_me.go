package company

import (
	"context"

	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateCompanyMe(ctx context.Context, req *v1.PatchUpdateCompanyMeReq) (res *v1.PatchUpdateCompanyMeRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	in := &v1.PatchUpdateCompanyByIDReq{
		Name: req.Name,
		Description: req.Description,
		Industry: req.Industry,
		Tag: req.Tag,
		Address: req.Address,
		Website: req.Website,
		City: req.City,
		Size: req.Size,
		Contact: req.Contact,
	}
	
	err = service.Company().PatchUpdateCompanyByID(ctx, in, *tokenData.CompanyID)

	return
}
