package user

import (
	"context"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteProfilePicture(ctx context.Context, req *v1.DeleteProfilePictureReq) (res *v1.DeleteProfilePictureRes, err error) {
	err = service.User().DeleteProfilePicture(ctx)

	return
}
