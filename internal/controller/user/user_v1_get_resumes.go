package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetResumes(ctx context.Context, req *v1.GetResumesReq) (res *v1.GetResumesRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	resumes, err := service.User().GetResumesByUserID(ctx, *tokenData.UserID)

	res = &v1.GetResumesRes{
		Resumes: resumes,
	}

	return
}
