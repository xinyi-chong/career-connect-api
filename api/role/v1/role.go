package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateRoleReq struct {
	g.Meta   `path:"/" method:"post" tags:"Role" summary:"Create Role"`
	Name		    string 	 `json:"name"       v:"required"`
}

type PostCreateRoleRes struct{
	Id 				*string				`json:"id"`
}

type GetRoleByIDReq struct {
	g.Meta   `path:"/:role_id" method:"get" tags:"Role" summary:"Get Role By ID"`
}

type GetRoleByIDRes struct {
	Role *entity.Role	`json:"Role"`
}

type PatchUpdateRoleByIDReq struct {
	g.Meta   `path:"/:role_id" method:"patch" tags:"Role" summary:"Update Role By ID"`
	Name		    string 	 `json:"name"       v:"required"`
	Status 			string   `json:"status"     v:"required"`
}

type PatchUpdateRoleByIDRes struct {}

type DeleteRoleByIDReq struct {
	g.Meta   `path:"/:role_id" method:"delete" tags:"Role" summary:"Delete Role By ID"`
}

type DeleteRoleByIDRes struct {}
