package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCertificateFirstCert(ctx context.Context, req *v1.GetCertificateFirstCertReq) (res *v1.GetCertificateFirstCertRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	certificates, err := service.User().GetCertificateFirstCert(ctx, tokenData.AccountID)

	res = &v1.GetCertificateFirstCertRes{
		FirstCertCertificates: certificates,
	}

	return
}
