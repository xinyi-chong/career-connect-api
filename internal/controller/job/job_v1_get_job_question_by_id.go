package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetJobQuestionByID(ctx context.Context, req *v1.GetJobQuestionByIDReq) (res *v1.GetJobQuestionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobQuestionID := r.GetRouter("question_id").String()

	jobQuestion, err := service.Job().GetJobQuestionByID(ctx, jobQuestionID)

	res = &v1.GetJobQuestionByIDRes{
		JobQuestion: jobQuestion,
	}

	return
}
