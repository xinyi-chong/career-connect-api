package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteJobQuestionByID(ctx context.Context, req *v1.DeleteJobQuestionByIDReq) (res *v1.DeleteJobQuestionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobQuestionID := r.GetRouter("question_id").String()

	err = service.Job().DeleteJobQuestionByID(ctx, jobQuestionID)

	return
}
