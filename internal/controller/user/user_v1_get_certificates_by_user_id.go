package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetCertificatesByUserID(ctx context.Context, req *v1.GetCertificatesByUserIDReq) (res *v1.GetCertificatesByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	certificates, err := service.User().GetCertificatesByUserID(ctx, userID);

	res = &v1.GetCertificatesByUserIDRes{
		Certificates: certificates,
	}

	return
}
