// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"gf_demo/api/role/v1"
)

type IRoleV1 interface {
	PostCreateRole(ctx context.Context, req *v1.PostCreateRoleReq) (res *v1.PostCreateRoleRes, err error)
	GetRoleByID(ctx context.Context, req *v1.GetRoleByIDReq) (res *v1.GetRoleByIDRes, err error)
	PatchUpdateRoleByID(ctx context.Context, req *v1.PatchUpdateRoleByIDReq) (res *v1.PatchUpdateRoleByIDRes, err error)
	DeleteRoleByID(ctx context.Context, req *v1.DeleteRoleByIDReq) (res *v1.DeleteRoleByIDRes, err error)
}
