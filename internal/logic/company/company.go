package company

import (
	"context"
	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sCompany struct{}

func init() {
	service.RegisterCompany(New())
}

func New() *sCompany {
	return &sCompany{}
}

func (s *sCompany) GetCompanyByID(ctx context.Context, companyID string) (*entity.Company, error) {
	cacheCompany := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID, &entity.Company{})
	if cacheCompany != nil {
		return cacheCompany.(*entity.Company), nil
	}

	var company *entity.Company
	err := dao.Company.Ctx(ctx).With(
		entity.Media{},
		entity.CompanySubscription{},
		entity.CompanyAccounts{},
	).Where(do.Company{
		Id: companyID,
	}).Scan(&company)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Company By Company ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Company By Company ID" + err.Error())
		return nil, err
	}
	
	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_COMPANY_ID + companyID, company)

	return company, nil
}

func (s *sCompany) GetCompanyByAccountID(ctx context.Context, accountID string) (*entity.Company, error) {
	var company *entity.Company
	err := dao.Company.Ctx(ctx).With(
		entity.Media{},
	).Where(do.Company{
		AccountId: accountID,
	}).Scan(&company)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Company By Account ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Company By Account ID" + err.Error())
	}

	return company, err
}

func (s *sCompany) PostCreateCompany(ctx context.Context, req *v1.PostCreateCompanyReq, accountID string) (companyID *string, err error) {
	id := uuid.New().String()
	
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Media (Logo)
		var mediaID *string
		if req.Logo != nil {
			if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				mediaID, err = service.Media().CreateMedia(ctx, req.Logo, accountID)
				return err
			}); err != nil {
				return err
			}
		}

		// Insert Company
		if err := dao.Company.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err := dao.Company.Ctx(ctx).Data(do.Company{
				Id: id,
				AccountId: accountID,
				Name: req.Name,
				Description: req.Description,
				Industry: req.Industry,
				Tag: req.Tag,
				Address: req.Address,
				Website: req.Website,
				City: req.City,
				Size: req.Size,
				Contact: req.Contact,
				LogoId: mediaID,
			}).Insert()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Create Company: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Company: " + err.Error())
			return err
		}

		g.Log().Info(ctx, consts.SUCCESS_CREATE, "company with id: ", id)
		return nil
	})

	return &id, err
}

// Update Company By ID (Without Logo)
func (s *sCompany) PatchUpdateCompanyByID(ctx context.Context, req *v1.PatchUpdateCompanyByIDReq, companyID string) (error) {
	_, err := dao.Company.Ctx(ctx).Data(do.Company{
		Name: req.Name,
		Description: req.Description,
		Industry: req.Industry,
		Tag: req.Tag,
		Address: req.Address,
		Website: req.Website,
		City: req.City,
		Size: req.Size,
		Contact: req.Contact,
	}).Where(do.Company{
		Id: companyID,
	}).Update()
	if err != nil {
		g.Log().Error(ctx, "Failed to Update Company By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Company By ID: " + err.Error())
		return err
	}

	service.Session().RemoveCompanySession(ctx, &companyID)

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "company")
	return nil
}

func (s *sCompany) DeleteCompanyMe(ctx context.Context) (err error) {	
	sessionData, err := service.Session().GetSessionDataFromCtx(ctx)
	if err != nil {
		return
	}

	logoID := sessionData.Company.LogoId
	companyID := sessionData.Company.Id

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Delete Account
		if err = dao.Account.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = service.Account().DeleteAccountByID(ctx, sessionData.Account.Id)
			return err
		}); err != nil {
			return err
		}

		// Delete Media (Resume & ProfilePicture)
		if err = dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// Delete Logo Media
			err = service.Media().DeleteMediaByID(ctx, logoID);
			return err
		}); err != nil {
			return err
		}

		g.Log().Info(ctx, consts.SUCCESS_DELETE, "Company Me")
		return nil
	})

	if err != nil {
		return err
	}

	// Remove Caches
	companyAccounts, err := service.Company().GetCompanyAccountsByCompanyID(ctx, companyID)
	for _, companyAccount := range companyAccounts {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + companyAccount.UserId)
	}
	cacheKeys := []string{
		consts.CACHE_COMPANY_ID + companyID,
		consts.CACHE_JOBS,
		consts.CACHE_JOBS_BY_COMPANY_ID + companyID,
		consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + sessionData.Account.Id,
		consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID + companyID,
	}
	service.Cache().RemoveMulCachesWithPrefix(ctx, cacheKeys)

	// Remove Company Session
	service.Session().RemoveSession(ctx)
	
	return nil
}