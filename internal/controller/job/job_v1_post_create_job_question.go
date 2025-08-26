package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateJobQuestion(ctx context.Context, req *v1.PostCreateJobQuestionReq) (res *v1.PostCreateJobQuestionRes, err error) {
	id, err := service.Job().PostCreateJobQuestion(ctx, req)

	res = &v1.PostCreateJobQuestionRes{
		Id: *id,
	}

	return
}
