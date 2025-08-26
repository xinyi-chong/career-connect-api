package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateUser(ctx context.Context, req *v1.PostCreateUserReq) (res *v1.PostCreateUserRes, err error) {
	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	} else if tokenData.UserID != nil {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "User already exists")
		return
	}

	id, err := service.User().PostCreateUser(ctx, req, tokenData.AccountID);

	res = &v1.PostCreateUserRes{
		Id: *id,
	}

	return
}
