package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetSkillsByUserID(ctx context.Context, req *v1.GetSkillsByUserIDReq) (res *v1.GetSkillsByUserIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	userID := r.GetRouter("user_id").String()

	skills, err := service.User().GetSkillsByUserID(ctx, userID);

	res = &v1.GetSkillsByUserIDRes{
		Skills: skills,
	}

	return
}
