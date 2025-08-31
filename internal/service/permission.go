// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/permission/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
)

type (
	IPermission interface {
		GetPermissionByID(ctx context.Context, permissionID string) (*entity.Permission, error)
		GetPermissionsByRoleID(ctx context.Context, roleID string) ([]*entity.Permission, error)
		GetPermissionsByFeatureID(ctx context.Context, featureID string) ([]*entity.Permission, error)
		GetPermissionByRoleIDFeatureID(ctx context.Context, roleID string, featureID string) (*entity.Permission, error)
		GetCompanyPermissionsByUserID(ctx context.Context, userID string) ([]*model.Permission, error)
		PostCreatePermission(ctx context.Context, req *v1.PostCreatePermissionReq) (*string, error)
		PatchUpdatePermissionByRoleIDFeatureID(ctx context.Context, req *v1.PatchUpdatePermissionByRoleIDFeatureIDReq, roleID string, featureID string) error
		DeletePermissionByID(ctx context.Context, permissionID string) error
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
