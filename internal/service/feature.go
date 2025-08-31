// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/feature/v1"
	"gf_demo/internal/model/entity"
)

type (
	IFeature interface {
		GetFeatureByID(ctx context.Context, featureID string) (*entity.Feature, error)
		PostCreateFeature(ctx context.Context, req *v1.PostCreateFeatureReq) (*string, error)
		PatchUpdateFeatureByID(ctx context.Context, req *v1.PatchUpdateFeatureByIDReq, featureID string) error
		DeleteFeatureByID(ctx context.Context, featureID string) error
	}
)

var (
	localFeature IFeature
)

func Feature() IFeature {
	if localFeature == nil {
		panic("implement not found for interface IFeature, forgot register?")
	}
	return localFeature
}

func RegisterFeature(i IFeature) {
	localFeature = i
}
