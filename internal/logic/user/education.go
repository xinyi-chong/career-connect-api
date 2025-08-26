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

func (s *sUser) GetEducationsByUserID(ctx context.Context, userID string) ([]*entity.Education, error) {
	var educations []*entity.Education
	err := dao.Education.Ctx(ctx).WithAll().Where(do.Education{
		UserId: userID,
	}).Scan(&educations)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Educations By User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Educations By User ID: " + err.Error())
	}

	return educations, err
}

func (s *sUser) GetEducationByID(ctx context.Context, educationID string) (*entity.Education, error) {
	var education *entity.Education
	err := dao.Education.Ctx(ctx).WithAll().Where(do.Education{
		Id: educationID,
	}).Scan(&education)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Education by ID: ", educationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Education By ID: " + err.Error())
	}

	return education, err
}

func (s *sUser) PostCreateEducation(ctx context.Context, req *v1.PostCreateEducationReq, userID string) (*string, error) {
	instituteString := req.InstituteString
	if req.InstituteID != nil && req.InstituteString == nil {
		company, err := service.Company().GetCompanyByID(ctx, *req.InstituteID)
		if err != nil {
			return nil, err
		} else if company == nil {
			g.Log().Error(ctx, "Company Not Found: ", req.InstituteID)
			err = gerror.NewCode(gcode.CodeNotFound, "Company Not Found")
			return nil, err
		}
		instituteString = &company.Name
	}

	educationID := uuid.New().String()
	_, err := dao.Education.Ctx(ctx).Data(do.Education{
		Id: educationID,
		UserId: userID,
		StartDate: &req.StartDate,
		EndDate: req.EndDate,
		InstituteId: req.InstituteID,
		InstituteString: instituteString,
		Level: req.Level,
		Programme: req.Programme,
		Description: req.Description,
	}).OmitEmptyData().Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Education: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Education: " + err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Education", educationID)
	return &educationID, nil
}

func (s *sUser) PatchUpdateEducationByID(ctx context.Context, req *v1.PatchUpdateEducationByIDReq, educationID string, userID string) (error) {	
	_, err := dao.Education.Ctx(ctx).Data(do.Education{
		StartDate: req.StartDate,
		EndDate: req.EndDate,
		InstituteId: req.InstituteID,
		InstituteString: req.InstituteString,
		Level: req.Level,
		Programme: req.Programme,
		Description: req.Description,
	}).Where(do.Education{
		Id: educationID,
		UserId: userID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Education By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Education By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Education", educationID)
	return nil
}

func (s *sUser) DeleteEducationByID(ctx context.Context, educationID string, userID string) error {
	_, err := dao.Education.Ctx(ctx).Where(do.Education{
		Id: educationID,
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Education by ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Education By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Education: ", educationID)
	return nil
}