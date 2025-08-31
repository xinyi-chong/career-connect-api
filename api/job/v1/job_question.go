package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateJobQuestionReq struct {
	g.Meta `path:"/question" tags:"Job" method:"post" summary:"Create Job Question"`
	Question      string      `json:"question"       v:"required"`
	JobID      	  string      `json:"job_id"         v:"required"`
}

type PostCreateJobQuestionRes struct {
	Id					string  		`json:"id"`
}

type GetJobQuestionByIDReq struct {
	g.Meta `path:"/question/:question_id" tags:"Job" method:"get" summary:"Get Job Question By ID"`
}

type GetJobQuestionByIDRes struct {
	JobQuestion *entity.JobQuestion `json:"job_question"`
}

type PatchUpdateJobQuestionByIDReq struct {
	g.Meta `path:"/question/:question_id" tags:"Job" method:"patch" summary:"Update Job Question By ID"`
	Question      string      `json:"question"       v:"required"`
}

type PatchUpdateJobQuestionByIDRes struct {
}

type DeleteJobQuestionByIDReq struct {
	g.Meta `path:"/question/:question_id" tags:"Job" method:"delete" summary:"Delete Job Question By ID"`
}

type DeleteJobQuestionByIDRes struct {
}

