package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCertificates(ctx context.Context, req *v1.GetCertificatesReq) (res *v1.GetCertificatesRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	certificates, err := service.User().GetCertificatesByUserID(ctx, *tokenData.UserID)

	res = &v1.GetCertificatesRes{
		Certificates: certificates,
	}

	return
}
