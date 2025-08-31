package v1

import (
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreatePermissionReq struct {
	g.Meta   `path:"/" method:"post" tags:"Permission" summary:"Create Permission"`
	RoleID		    string 	 `json:"role_id"       v:"required"`
	FeatureID		  string 	 `json:"feature_id"    v:"required"`
	Allow		      bool 	   `json:"allow"         v:"required"`
}

type PostCreatePermissionRes struct{
	Id 				*string				`json:"id"`
}

type GetPermissionByIDReq struct {
	g.Meta   `path:"/:permission_id" method:"get" tags:"Permission" summary:"Get Permission By ID"`
}

type GetPermissionByIDRes struct {
	Permission *entity.Permission	`json:"permission"`
}

type GetCompanyPermissionsByUserIDReq struct {
	g.Meta   `path:"/user/:user_id" method:"get" tags:"Permission" summary:"Get Permission By User ID"`
}

type GetCompanyPermissionsByUserIDRes struct {
	Permissions []*model.Permission	`json:"permissions"`
}

type GetPermissionsByRoleIDReq struct {
	g.Meta   `path:"/role/:role_id" method:"get" tags:"Permission" summary:"Get Permission By Role ID"`
}

type GetPermissionsByRoleIDRes struct {
	Permissions []*entity.Permission	`json:"permissions"`
}

type GetPermissionByRoleIDFeatureIDReq struct {
	g.Meta   `path:"/role/:role_id/feature/:feature_id" method:"get" tags:"Permission" summary:"Get Permission By Role ID & Feature ID"`
}

type GetPermissionByRoleIDFeatureIDRes struct {
	Permission *entity.Permission	`json:"permission"`
}

type PatchUpdatePermissionByRoleIDFeatureIDReq struct {
	g.Meta   `path:"/role/:role_id/feature/:feature_id" method:"patch" tags:"Permission" summary:"Update Permission By Role ID & Feature ID"`
	Allow		      bool 	   `json:"allow"         v:"required"`
}

type PatchUpdatePermissionByRoleIDFeatureIDRes struct {}

type DeletePermissionByIDReq struct {
	g.Meta   `path:"/:permission_id" method:"delete" tags:"Permission" summary:"Delete Permission By ID"`
}

type DeletePermissionByIDRes struct {}
