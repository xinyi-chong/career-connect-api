package feature

import (
	"context"

	v1 "gf_demo/api/feature/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
)

type sFeature struct{}

func init() {
	service.RegisterFeature(New())
}

func New() *sFeature {
	return &sFeature{}
}


func (s *sFeature) GetFeatureByID(ctx context.Context, featureID string) (*entity.Feature, error) {
	cacheFeature := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_FEATURE_ID + featureID, &entity.Feature{})
	if cacheFeature != nil {
		return cacheFeature.(*entity.Feature), nil
	}

	var feature *entity.Feature
	err := dao.Feature.Ctx(ctx).Where(do.Feature{
		Id: featureID,
	}).Scan(&feature)

	if err != nil {
		g.Log().Error("Failed to Get Feature By ID:", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Feature By ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_FEATURE_ID + featureID, feature)
	return feature, err
}

func (s *sFeature) PostCreateFeature(ctx context.Context, req *v1.PostCreateFeatureReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Feature.Ctx(ctx).Data(do.Feature{
		Id: id,
		Name: req.Name,
	}).Insert()

	if err != nil {
		g.Log().Error("Failed to Create Feature", err) 
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Feature: " + err.Error())
	}

	g.Log().Info(consts.SUCCESS_CREATE, "Feature: ", id)
	return &id, err
}

func (s *sFeature) PatchUpdateFeatureByID(ctx context.Context, req *v1.PatchUpdateFeatureByIDReq, featureID string) (error) {
	_, err := dao.Feature.Ctx(ctx).Data(do.Feature{
		Name: req.Name,
	}).Where(do.Feature{
		Id: featureID,
	}).Update()
	if err != nil {
		g.Log().Error("Failed to Update Feature By ID", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Feature By ID: " + err.Error())
		return err
	}

	// Remove Caches
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_FEATURE_ID + featureID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID)

	g.Log().Info(consts.SUCCESS_UPDATE, "Feature By ID", featureID) 
	return nil
}

func (s *sFeature) DeleteFeatureByID(ctx context.Context, featureID string) (error) {
	_, err := dao.Feature.Ctx(ctx).Where(do.Feature{
		Id: featureID,
	}).Delete()
	if err != nil {
		g.Log().Error("Failed to Delete Feature By ID", featureID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Feature By ID: " + err.Error())
		return err
	}

	// Remove Permissions Cache
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_FEATURE_ID + featureID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_PERMISSIONS_BY_ROLE_ID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_COMPANY_ID)
	service.Cache().RemoveCacheMatchedPattern(ctx, consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID)
	
	// Remove All Company Accounts Sessions
	var companyaccounts []*entity.CompanyAccounts
	err = dao.Companyaccounts.Ctx(ctx).Fields("user_id").Distinct().Scan(&companyaccounts)
	if err == nil {
		service.Session().RemoveCompanyAccountsSessions(ctx, companyaccounts)
	}

	g.Log().Info(consts.SUCCESS_DELETE, "Feature By ID", featureID)
	return nil
}