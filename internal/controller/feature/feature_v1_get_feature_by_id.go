package feature

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/feature/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) GetFeatureByID(ctx context.Context, req *v1.GetFeatureByIDReq) (res *v1.GetFeatureByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	featureID := r.GetRouter("feature_id").String()
 
	feature, err := service.Feature().GetFeatureByID(ctx, featureID);
	if err != nil {
		return
	}

	res = &v1.GetFeatureByIDRes{
		Feature: feature,
	}

	return
}
