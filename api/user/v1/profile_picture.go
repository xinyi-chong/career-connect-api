package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Profile Picture
type PostCreateProfilePictureReq struct {
	g.Meta `path:"/profile_picture" tags:"User" method:"post" summary:"Create Profile Picture"`
	ProfilePicture  *ghttp.UploadFile  		`json:"profile_picture" v:"required"`
}

type PostCreateProfilePictureRes struct {
	Id     string      `json:"id"`
}

type PatchUpdateProfilePictureReq struct {
	g.Meta `path:"/profile_picture" tags:"User" method:"patch" summary:"Update Profile Picture"`
	ProfilePicture  *ghttp.UploadFile  		`json:"profile_picture" v:"required"`
}

type PatchUpdateProfilePictureRes struct {
}

type DeleteProfilePictureReq struct {
	g.Meta `path:"/profile_picture" tags:"User" method:"delete" summary:"Delete Profile Picture"`
}

type DeleteProfilePictureRes struct {
}