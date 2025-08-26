package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateCertificate(ctx context.Context, req *v1.PostCreateCertificateReq) (res *v1.PostCreateCertificateRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}
	
	id, err := service.User().PostCreateCertificate(ctx, req, *tokenData.UserID)
	
	res = &v1.PostCreateCertificateRes{
		Id: *id,
	}

	return
}
