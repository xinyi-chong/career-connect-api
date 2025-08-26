package feature

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "gf_demo/api/feature/v1"
	"gf_demo/internal/service"
)

func (c *ControllerV1) DeleteFeatureByID(ctx context.Context, req *v1.DeleteFeatureByIDReq) (res *v1.DeleteFeatureByIDRes, err error) {
	r := g.RequestFromCtx(ctx)
	featureID := r.GetRouter("feature_id").String()
 
	err = service.Feature().DeleteFeatureByID(ctx, featureID);

	return
}
