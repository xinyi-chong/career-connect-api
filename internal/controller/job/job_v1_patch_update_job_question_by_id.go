package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) PatchUpdateJobQuestionByID(ctx context.Context, req *v1.PatchUpdateJobQuestionByIDReq) (res *v1.PatchUpdateJobQuestionByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	jobQuestionID := r.GetRouter("question_id").String()

	err = service.Job().PatchUpdateJobQuestionByID(ctx, req, jobQuestionID)

	return
}
