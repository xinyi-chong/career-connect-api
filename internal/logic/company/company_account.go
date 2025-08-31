package company

import (
	"context"
	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sCompany) GetCompanyAccountsByCompanyID(ctx context.Context, companyID string) ([]*entity.CompanyAccounts, error) {
	cacheCompanyAccounts := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID + companyID, &[]*entity.CompanyAccounts{})
	if cacheCompanyAccounts != nil {
		return *cacheCompanyAccounts.(*[]*entity.CompanyAccounts), nil
	}
	
	var companyAccounts []*entity.CompanyAccounts
	err := dao.Companyaccounts.Ctx(ctx).With(
		entity.Role{},
		entity.Role{}.Permissions,
		entity.Feature{},
	).Where(do.CompanyAccounts{
		CompanyId: companyID,
	}).Scan(&companyAccounts)
	if err != nil {
		g.Log().Error(ctx, "Failed to get company accounts by company ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to get company accounts by company ID" + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID + companyID, companyAccounts)

	return companyAccounts, err
}

func (s *sCompany) GetCompanyAccountsByUserID(ctx context.Context, userID string) ([]*entity.CompanyAccounts, error) {
	cacheCompanyAccounts := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + userID, &[]*entity.CompanyAccounts{})
	if cacheCompanyAccounts != nil {
		return *cacheCompanyAccounts.(*[]*entity.CompanyAccounts), nil
	}

	var companyAccounts []*entity.CompanyAccounts
	err := dao.Companyaccounts.Ctx(ctx).With(
		entity.Role{},
		entity.Role{}.Permissions,
		entity.Feature{},
	).Where(do.CompanyAccounts{
		UserId: userID,
	}).Scan(&companyAccounts)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get CompanyAccounts By User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get CompanyAccounts By User ID" + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + userID, companyAccounts)
		
	return companyAccounts, err
}

func (s *sCompany) GetCompanyAccountsByRoleID(ctx context.Context, roleID string) ([]*entity.CompanyAccounts, error) {
	var companyaccounts []*entity.CompanyAccounts
	err := dao.Companyaccounts.Ctx(ctx).With(
		entity.User{},
		entity.Company{},
	).Where(do.CompanyAccounts{
		RoleId: roleID,
	}).Scan(&companyaccounts)

	if err != nil {
		g.Log().Error(ctx, "Failed to get company accounts by Role ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Company Accounts by Role ID" + err.Error())
		return nil, err
	}
		
	return companyaccounts, nil
}

func (s *sCompany) PostCreateCompanyAccount(ctx context.Context, req *v1.PostCreateCompanyAccountReq, companyID string) (error) {
	_, err := dao.Companyaccounts.Ctx(ctx).Data(do.CompanyAccounts{
		CompanyId: companyID,
		UserId: req.UserID,
		RoleId: req.RoleID,
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "Failed to Create CompanyAccounts", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create CompanyAccounts" + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "company accounts")
	return nil
}

func (s *sCompany) PostCreateCompanyAccountByCompanyID(ctx context.Context, req *v1.PostCreateCompanyAccountByCompanyIDReq, companyID string) (error) {
	_, err := dao.Companyaccounts.Ctx(ctx).Data(do.CompanyAccounts{
		CompanyId: companyID,
		UserId: req.UserID,
		RoleId: req.RoleID,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create CompanyAccounts By Company ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create CompanyAccounts By Company ID" + err.Error())
		return err
	}

	s.RemoveCompanyAccountCache(ctx, companyID, req.UserID)

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "company accounts by Company ID")
	return nil
}

func (s *sCompany) PatchUpdateCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByCompanyIDUserIDReq, companyID string, userID string) (error) {
	_, err := dao.Companyaccounts.Ctx(ctx).Data(do.CompanyAccounts{
		RoleId: req.RoleID,
	}).Where(do.CompanyAccounts{
		CompanyId: companyID,
		UserId: userID,
	}).Update()
	if err != nil {
		g.Log().Error(ctx, "Failed to Update CompanyAccounts By Company ID & User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update CompanyAccounts By Company ID & User ID" + err.Error())
		return err
	}

	s.RemoveCompanyAccountCache(ctx, companyID, userID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "company accounts by Company ID & User ID")
	return nil
}

func (s *sCompany) DeleteCompanyAccountByCompanyIDUserID(ctx context.Context, companyID string, userID string) (error) {
	_, err := dao.Companyaccounts.Ctx(ctx).Where(do.CompanyAccounts{
		CompanyId: companyID,
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete CompanyAccounts By Account ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete CompanyAccounts By Account ID" + err.Error())
		return err
	}

	s.RemoveCompanyAccountCache(ctx, companyID, userID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "company accounts by Account ID")
	return nil
}

func (s *sCompany) RemoveCompanyAccountCache(ctx context.Context, companyID string, userID string) {
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID + companyID)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + userID)
}

func (s *sCompany) RemoveCompanyAccountsCaches(ctx context.Context, companyaccounts []*entity.CompanyAccounts) {
	for _, companyaccount := range companyaccounts {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID + companyaccount.CompanyId)
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + companyaccount.UserId)
	}
}