package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCertificateByID(ctx context.Context, req *v1.GetCertificateByIDReq) (res *v1.GetCertificateByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	certificateID := r.GetRouter("certificate_id").String()

	certificate, err := service.User().GetCertificateByID(ctx, certificateID);

	res = &v1.GetCertificateByIDRes{
		Certificate: certificate,
	}

	return
}
