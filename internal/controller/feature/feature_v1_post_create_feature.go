package feature

import (
	"context"

	v1 "gf_demo/api/feature/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) PostCreateFeature(ctx context.Context, req *v1.PostCreateFeatureReq) (res *v1.PostCreateFeatureRes, err error) {
	featureID, err := service.Feature().PostCreateFeature(ctx, req);

	if err != nil {
		return
	}

	res = &v1.PostCreateFeatureRes{
		Id: featureID,
	}

	return
}
