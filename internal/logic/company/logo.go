package company

import (
	"context"
	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sCompany) PostCreateLogoByCompanyID(ctx context.Context, req *v1.PostCreateLogoByCompanyIDReq, companyID string) (error) {
	company, err := s.GetCompanyByID(ctx, companyID)
	if err != nil {
		return err
	}

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Media (Logo)
		var mediaID *string
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			mediaID, err = service.Media().CreateMedia(ctx, &req.Logo, company.AccountId)
			return err
		}); err != nil {
			return err
		}

		// Update Company's LogoId
		if err := dao.Company.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.Company.Ctx(ctx).Data(do.Company{
				LogoId: mediaID,
			}).Where(do.Company{
				Id: companyID,
			}).Update()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Update LogoId by company ID: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update LogoId by company ID: " + err.Error())
			return err
		}
		g.Log().Info(ctx, consts.SUCCESS_CREATE, "Logo by company ID: ", mediaID)
		return nil
	})

	service.Session().RemoveCompanySession(ctx, &companyID)

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID)

	return err
}

func (s *sCompany) PatchUpdateLogoByCompanyID(ctx context.Context, req *v1.PatchUpdateLogoByCompanyIDReq, companyID string) (error) {
	company, err := service.Company().GetCompanyByID(ctx, companyID)
	if err != nil {
		return err
	}
	
	err = service.Media().UpdateMediaByID(ctx, company.LogoId, &req.Logo, company.AccountId)
	if err != nil {
		return err
	}

	service.Session().RemoveCompanySession(ctx, &companyID)

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID)

	return nil
}

func (s *sCompany) DeleteLogoByCompanyID(ctx context.Context, companyID string) (error) {
	company, err := s.GetCompanyByID(ctx, companyID)
	if err != nil {
		return err
	}
	
	err = service.Media().DeleteMediaByID(ctx, company.LogoId)
	if err != nil {
		return err
	}

	service.Session().RemoveCompanySession(ctx, &companyID)

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "logo by company ID")
	return nil
}