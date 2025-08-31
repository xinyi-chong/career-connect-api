package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateUserByID(ctx context.Context, req *v1.PatchUpdateUserByIDReq) (res *v1.PatchUpdateUserByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	err = service.User().PatchUpdateUserByID(ctx, req, userID)

	return
}
