package user

import (
	"context"
	v1 "gf_demo/api/user/v1"
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

func (s *sUser) GetExperiencesByUserID(ctx context.Context, userID string) ([]*entity.Experience, error) {
	var experiences []*entity.Experience
	err := dao.Experience.Ctx(ctx).With(
		entity.Company{},
		entity.Company{}.Logo,
	).Where(do.Experience{
		UserId: userID,
	}).Scan(&experiences)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get User Experiences By User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get User Experiences By User ID: " + err.Error())
	}

	return experiences, err
}

func (s *sUser) GetExperienceByID(ctx context.Context, experienceID string) (*entity.Experience, error) {
	var experience *entity.Experience
	err := dao.Experience.Ctx(ctx).With(
		entity.Company{},
		entity.Company{}.Logo,
	).Where(do.Experience{
		Id: experienceID,
	}).Scan(&experience)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get User Experience By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get User Experience By ID: " + err.Error())
	}

	return experience, err
}

func (s *sUser) PostCreateExperience(ctx context.Context, req *v1.PostCreateExperienceReq, userID string) (*string, error) {
	companyString := req.CompanyString
	if req.CompanyID != nil && (req.CompanyString == nil || *req.CompanyString == "") {
		company, err := service.Company().GetCompanyByID(ctx, *req.CompanyID)
		if err != nil {
			return nil, err
		} else if company == nil {
			err = gerror.NewCode(gcode.CodeNotFound, "Company Not Found")
			return nil, err
		}
		companyString = &company.Name
	}

	experienceID := uuid.New().String()
	_, err := dao.Experience.Ctx(ctx).Data(do.Experience{
		Id: experienceID,
		UserId: userID,
		StartDate: &req.StartDate,
		EndDate: req.EndDate,
		IsPresent: req.EndDate == nil,
		Description: &req.Description,
		Title: req.Title,
		CompanyId: req.CompanyID,
		CompanyString: companyString,
	}).Insert()
	
	if err != nil {
		g.Log().Error(ctx, "Failed to Create User Experience: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create User Experience: " + err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)
	
	g.Log().Info(ctx, consts.SUCCESS_CREATE, "User Experience: ", experienceID)
	return &experienceID, nil
}

func (s *sUser) PatchUpdateExperienceByID(ctx context.Context, req *v1.PatchUpdateExperienceByIDReq, experienceID string, userID string) error {
	_, err := dao.Experience.Ctx(ctx).Data(do.Experience{
		StartDate: req.StartDate,
		EndDate: req.EndDate,
		IsPresent: req.IsPresent,
		Description: &req.Description,
		Title: req.Title,
		CompanyId: req.CompanyID,
		CompanyString: req.CompanyString,
	}).Where(do.Experience{
		Id: experienceID,
		UserId: userID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update User Experience By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update User Experience By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "User Experience By ID: ", experienceID)
	return nil
}

func (s *sUser) DeleteExperienceByID(ctx context.Context, experienceID string, userID string) (err error) {
	_, err = dao.Experience.Ctx(ctx).Where(do.Experience{
		Id: experienceID,
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete User Experience By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete User Experience By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)
	
	g.Log().Info(ctx, consts.SUCCESS_DELETE, "User Experience By ID: ", experienceID)
	return nil
}