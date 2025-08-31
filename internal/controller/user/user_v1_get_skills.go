package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetSkills(ctx context.Context, req *v1.GetSkillsReq) (res *v1.GetSkillsRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}

	skills, err := service.User().GetSkillsByUserID(ctx, *tokenData.UserID)

	res = &v1.GetSkillsRes{
		Skills: skills,
	}

	return
}
