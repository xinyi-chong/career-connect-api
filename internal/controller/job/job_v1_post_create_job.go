package job

import (
	"context"

	v1 "gf_demo/api/job/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/model"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateJob(ctx context.Context, req *v1.PostCreateJobReq) (res *v1.PostCreateJobRes, err error) {
	var createdBy string
	var createdByType string

	tokenData, err := service.Token().GetTokenDataFromCtxVar(ctx)
	if err != nil {
		return
	}

	companyID := req.CompanyID
	createdByType = consts.COMPANY

	if companyID == nil && tokenData.CompanyID != nil {
		companyID = tokenData.CompanyID	
		createdBy = *tokenData.CompanyID
	} else if tokenData.UserID != nil {
		createdBy = *tokenData.UserID
		createdByType = consts.USER
	}
	
	in := model.PostCreateJobInput{
		Title: req.Title,
		CompanyID: *companyID,
		Tag: req.Tag,
		Description: req.Description,
		Level: req.Level,
		Salary: req.Salary,
		PostedAt: req.PostedAt,
		Location: req.Location,
		IsRemote: req.IsRemote,
		IsHybrid: req.IsHybrid,
		Expiry: req.Expiry,
		CreatedBy: createdBy,
		CreatedByType: createdByType,
		UpdatedBy: createdBy,
		UpdatedByType: createdByType,
	}

	id, err := service.Job().PostCreateJob(ctx, in)

	res = &v1.PostCreateJobRes{
		Id: id,
	}

	return
}
