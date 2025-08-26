package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateCertificateByID(ctx context.Context, req *v1.PatchUpdateCertificateByIDReq) (res *v1.PatchUpdateCertificateByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	certificateID := r.GetRouter("certificate_id").String()

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	err = service.User().PatchUpdateCertificateByID(ctx, req, certificateID, *tokenData.UserID)

	return
}
