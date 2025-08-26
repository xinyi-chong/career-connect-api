package suggestskill

import (
	"context"

	v1 "gf_demo/api/suggest_skill/v1"
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

type sSuggestSkill struct{}

func init() {
	service.RegisterSuggestSkill(New())
}

func New() *sSuggestSkill {
	return &sSuggestSkill{}
}

func (s *sSuggestSkill) GetSuggestSkillByID(ctx context.Context, suggestSkillID string) (*entity.SuggestSkill, error) {
	cacheSuggestSkill := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_SUGGEST_SKILL_ID + suggestSkillID, &entity.SuggestSkill{})
	if cacheSuggestSkill != nil {
		return cacheSuggestSkill.(*entity.SuggestSkill), nil
	}
	
	var suggestSkill *entity.SuggestSkill
	err := dao.Suggestskill.Ctx(ctx).Where(do.SuggestSkill{
		Id: suggestSkillID,
	}).Scan(&suggestSkill)

	if err != nil {
		g.Log().Error("Failed to Get Suggest Skill By ID:", suggestSkillID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Suggest Skill By ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_SUGGEST_SKILL_ID + suggestSkillID, suggestSkill)

	return suggestSkill, nil
}

func (s *sSuggestSkill) PostCreateSuggestSkill(ctx context.Context, req *v1.PostCreateSuggestSkillReq) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Suggestskill.Ctx(ctx).Data(do.SuggestSkill{
		Id: id,
		Name: req.Name,
		Category: req.Category,
	}).Insert()

	if err != nil {
		g.Log().Error("Failed to Create Suggest Skill", err) 
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Suggest Skill: " + err.Error())
		return nil, err
	}

	g.Log().Info(consts.SUCCESS_CREATE, "Suggest Skill: ", id)
	return &id, nil
}

func (s *sSuggestSkill) PatchUpdateSuggestSkillByID(ctx context.Context, req *v1.PatchUpdateSuggestSkillByIDReq, suggestSkillID string) (error) {
	_, err := dao.Suggestskill.Ctx(ctx).Data(do.SuggestSkill{
		Name: req.Name,
		Category: req.Category,
	}).Where(do.SuggestSkill{
		Id: suggestSkillID,
	}).Update()

	if err != nil {
		g.Log().Error("Failed to Update Suggest Skill By ID", suggestSkillID, err) 
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Suggest Skill By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_SUGGEST_SKILL_ID + suggestSkillID)

	g.Log().Info(consts.SUCCESS_UPDATE, "Suggest Skill By ID", suggestSkillID) 
	return nil
}

func (s *sSuggestSkill) DeleteSuggestSkillByID(ctx context.Context, suggestSkillID string) (error) {
	_, err := dao.Suggestskill.Ctx(ctx).Where(do.SuggestSkill{
		Id: suggestSkillID,
	}).Delete()

	if err != nil {
		g.Log().Error("Failed to Delete Suggest Skill By ID", suggestSkillID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Suggest Skill By ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_SUGGEST_SKILL_ID + suggestSkillID)

	g.Log().Info(consts.SUCCESS_DELETE, "Suggest Skill By ID", suggestSkillID)
	return nil
}