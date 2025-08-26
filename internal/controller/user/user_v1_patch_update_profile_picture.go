package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PatchUpdateProfilePicture(ctx context.Context, req *v1.PatchUpdateProfilePictureReq) (res *v1.PatchUpdateProfilePictureRes, err error) {
	err = service.User().PatchUpdateProfilePicture(ctx, req)

	return
}
