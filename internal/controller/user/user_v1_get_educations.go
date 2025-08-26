package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetEducations(ctx context.Context, req *v1.GetEducationsReq) (res *v1.GetEducationsRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	educations, err := service.User().GetEducationsByUserID(ctx, *tokenData.UserID);

	res = &v1.GetEducationsRes{
		Educations: educations,
	}

	return
}
