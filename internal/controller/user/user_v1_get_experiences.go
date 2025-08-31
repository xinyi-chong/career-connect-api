package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetExperiences(ctx context.Context, req *v1.GetExperiencesReq) (res *v1.GetExperiencesRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return nil, err
	}
	
	experiences, err := service.User().GetExperiencesByUserID(ctx, *tokenData.UserID);

	res = &v1.GetExperiencesRes{
		Experiences: experiences,
	}

	return
}
