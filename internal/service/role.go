// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/role/v1"
	"gf_demo/internal/model/entity"
)

type (
	IRole interface {
		GetRoleByID(ctx context.Context, roleID string) (*entity.Role, error)
		PostCreateRole(ctx context.Context, req *v1.PostCreateRoleReq) (*string, error)
		PatchUpdateRoleByID(ctx context.Context, req *v1.PatchUpdateRoleByIDReq, roleID string) error
		DeleteRoleByID(ctx context.Context, roleID string) error
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
