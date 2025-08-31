package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteCertificateByID(ctx context.Context, req *v1.DeleteCertificateByIDReq) (res *v1.DeleteCertificateByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	certificateID := r.GetRouter("certificate_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().DeleteCertificateByID(ctx, certificateID, *tokenData.UserID)

	return
}
