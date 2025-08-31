// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package feature

import (
	"context"

	"gf_demo/api/feature/v1"
)

type IFeatureV1 interface {
	PostCreateFeature(ctx context.Context, req *v1.PostCreateFeatureReq) (res *v1.PostCreateFeatureRes, err error)
	GetFeatureByID(ctx context.Context, req *v1.GetFeatureByIDReq) (res *v1.GetFeatureByIDRes, err error)
	PatchUpdateFeatureByID(ctx context.Context, req *v1.PatchUpdateFeatureByIDReq) (res *v1.PatchUpdateFeatureByIDRes, err error)
	DeleteFeatureByID(ctx context.Context, req *v1.DeleteFeatureByIDReq) (res *v1.DeleteFeatureByIDRes, err error)
}
