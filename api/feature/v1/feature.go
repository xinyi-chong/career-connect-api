package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateFeatureReq struct {
	g.Meta   `path:"/" method:"post" tags:"Feature" summary:"Create Feature"`
	Name		    string 	 `json:"name"       v:"required"`
}

type PostCreateFeatureRes struct{
	Id 				*string				`json:"id"`
}

type GetFeatureByIDReq struct {
	g.Meta   `path:"/:feature_id" method:"get" tags:"Feature" summary:"Get Feature By ID"`
}

type GetFeatureByIDRes struct {
	Feature *entity.Feature	`json:"Feature"`
}

type PatchUpdateFeatureByIDReq struct {
	g.Meta   `path:"/:feature_id" method:"patch" tags:"Feature" summary:"Update Feature By ID"`
	Name		    string 	 `json:"name"       v:"required"`
}

type PatchUpdateFeatureByIDRes struct {}

type DeleteFeatureByIDReq struct {
	g.Meta   `path:"/:feature_id" method:"delete" tags:"Feature" summary:"Delete Feature By ID"`
}

type DeleteFeatureByIDRes struct {}
