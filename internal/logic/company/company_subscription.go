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
	"github.com/google/uuid"
)

func (s *sCompany) GetCompanySubscriptionByCompanyID(ctx context.Context, companyID string) (*entity.CompanySubscription, error) {
	var companySubscription *entity.CompanySubscription
	err := dao.Companysubscription.Ctx(ctx).With(
		entity.CompanyPlan{},
	).Where(do.CompanySubscription{
		CompanyId: companyID,
	}).Scan(&companySubscription)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Company Subscription Me: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Company Subscription Me" + err.Error())
		return nil, err
	}

	return companySubscription, nil
}

func (s *sCompany) GetCompanySubscriptionByID(ctx context.Context, subscriptionID string) (*entity.CompanySubscription, error) {
	var companySubscription *entity.CompanySubscription
	err := dao.Companysubscription.Ctx(ctx).With(
		entity.CompanyPlan{},
	).Where(do.CompanySubscription{
		Id: subscriptionID,
	}).Scan(&companySubscription)
	
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Company Subscription By ID: ", subscriptionID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Company Subscription By ID" + err.Error())
		return nil, err
	}
	
	return companySubscription, nil
}

func (s *sCompany) PostCreateCompanySubscription(ctx context.Context, req *v1.PostCreateCompanySubscriptionReq, companyID string) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Companysubscription.Ctx(ctx).Data(do.CompanySubscription{
		Id: id,
		CompanyId: companyID,
		CompanyPlanId: req.CompanyPlanID,
		Expiry: req.Expiry,
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "Failed to Create Company Subscription: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Company Subscription" + err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companyID)

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "company subscription with id: ", id)
	return &id, err
}

func (s *sCompany) PatchUpdateCompanySubscriptionByID(ctx context.Context, req *v1.PatchUpdateCompanySubscriptionByIDReq, subscriptionID string) error {
	_, err := dao.Companysubscription.Ctx(ctx).Data(do.CompanySubscription{
		CompanyPlanId: req.CompanyPlanID,
		Expiry:        req.Expiry,
	}).Where(do.CompanySubscription{
		Id: subscriptionID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Company Subscription: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Company Subscription" + err.Error())
		return err
	}

	companySubscription, _ := service.Company().GetCompanySubscriptionByID(ctx, subscriptionID)
	if companySubscription != nil {
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_COMPANY_ID + companySubscription.CompanyId)
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Company Subscription")
	return err
}