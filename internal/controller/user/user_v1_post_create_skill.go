package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateSkill(ctx context.Context, req *v1.PostCreateSkillReq) (res *v1.PostCreateSkillRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	id, err := service.User().PostCreateSkill(ctx, req, *tokenData.UserID)

	res = &v1.PostCreateSkillRes{
		Id: *id,
	}

	return
}
