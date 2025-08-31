package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCertificateCertService(ctx context.Context, req *v1.GetCertificateCertServiceReq) (res *v1.GetCertificateCertServiceRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	certificates, err := service.User().GetCertificateCertService(ctx, tokenData.AccountID)

	res = &v1.GetCertificateCertServiceRes{
		CertServiceCertificates: certificates,
	}

	return
}
