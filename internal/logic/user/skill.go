package user

import (
	"context"
	v1 "gf_demo/api/user/v1"
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

func (s *sUser) GetSkillsByUserID(ctx context.Context, userID string) ([]*entity.Skill, error) {
	var skills []*entity.Skill
	err := dao.Skill.Ctx(ctx).Where(do.Skill{
		UserId: userID,
	}).Scan(&skills)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Skill By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Skill By ID: " + err.Error())
	}

	return skills, err
}

func (s *sUser) PostCreateSkill(ctx context.Context, req *v1.PostCreateSkillReq, userID string) (*string, error) {
	skillID := uuid.New().String()
	_, err := dao.Skill.Ctx(ctx).Data(do.Skill{
		Id: skillID,
		UserId: userID,
		Name: req.Name,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Skill: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Skill: " + err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Skill: ", skillID)
	return &skillID, err
}

func (s *sUser) PatchUpdateSkillByID(ctx context.Context, req *v1.PatchUpdateSkillByIDReq, skillID string, userID string) error {
	_, err := dao.Skill.Ctx(ctx).Data(do.Skill{
		Name: req.Name,
	}).Where(do.Skill{
		Id: skillID,
		UserId: userID,
	}).Update()

	if err != nil {
		g.Log().Info(ctx, "Failed to Update Skill: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Skill: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Skill: ", skillID)
	return nil
}

func (s *sUser) DeleteSkillByID(ctx context.Context, skillID string, userID string) error {
	_, err := dao.Skill.Ctx(ctx).Where(do.Skill{
		Id: skillID,
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Skill by ID: ", skillID, ", UserID: ", userID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Skill: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Skill: ", skillID)
	return err
}
