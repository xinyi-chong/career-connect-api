// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package permission

import (
	"context"

	"gf_demo/api/permission/v1"
)

type IPermissionV1 interface {
	PostCreatePermission(ctx context.Context, req *v1.PostCreatePermissionReq) (res *v1.PostCreatePermissionRes, err error)
	GetPermissionByID(ctx context.Context, req *v1.GetPermissionByIDReq) (res *v1.GetPermissionByIDRes, err error)
	GetCompanyPermissionsByUserID(ctx context.Context, req *v1.GetCompanyPermissionsByUserIDReq) (res *v1.GetCompanyPermissionsByUserIDRes, err error)
	GetPermissionsByRoleID(ctx context.Context, req *v1.GetPermissionsByRoleIDReq) (res *v1.GetPermissionsByRoleIDRes, err error)
	GetPermissionByRoleIDFeatureID(ctx context.Context, req *v1.GetPermissionByRoleIDFeatureIDReq) (res *v1.GetPermissionByRoleIDFeatureIDRes, err error)
	PatchUpdatePermissionByRoleIDFeatureID(ctx context.Context, req *v1.PatchUpdatePermissionByRoleIDFeatureIDReq) (res *v1.PatchUpdatePermissionByRoleIDFeatureIDRes, err error)
	DeletePermissionByID(ctx context.Context, req *v1.DeletePermissionByIDReq) (res *v1.DeletePermissionByIDRes, err error)
}
