package company

import (
	"context"
	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

func (s *sCompany) PostCreateCompanyPlan(ctx context.Context, req *v1.PostCreateCompanyPlanReq) (*string, error) {
	id := uuid.New().String()

	_, err := dao.Companyplan.Ctx(ctx).Data(do.CompanyPlan{
		Id:   id,
		Name: req.Name,
	}).Insert()
	if err != nil {
		g.Log().Error(ctx, "Failed to Create Company Plan: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Company Plan" + err.Error())
		return nil, err
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "company plan with id: ", id)
	return &id, nil
}

func (s *sCompany) PatchUpdateCompanyPlanByID(ctx context.Context, req *v1.PatchUpdateCompanyPlanByIDReq, companyPlanID string) (error) {
	_, err := dao.Companyplan.Ctx(ctx).Data(do.CompanyPlan{
		Name: req.Name,
	}).Where(do.CompanyPlan{
		Id: companyPlanID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Company Plan: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Company Plan" + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "company plan")
	return nil
}