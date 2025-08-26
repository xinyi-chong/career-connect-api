package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type GetUserMeReq struct {
	g.Meta `path:"/me" tags:"User" method:"get" summary:"Get User Me"`
}

type GetUserMeRes struct {
	User entity.User `json:"user"`
}

type GetUserByIDReq struct {
	g.Meta `path:"/:user_id" tags:"User" method:"get" summary:"Get User By ID"`
}

type GetUserByIDRes struct {
	User *entity.User `json:"user"`
}

type PostCreateUserReq struct {
	g.Meta `path:"/" tags:"User" method:"post" summary:"Create User"`
	Firstname       string      `json:"firstname"    v:"required"`
	Lastname        string      `json:"lastname"     v:"required"`
	Nationality  		string      `json:"nationality"  v:"required"`
	ProfilePicture  *ghttp.UploadFile      `json:"profile_picture"`
}

type PostCreateUserRes struct {
	Id					string  		`json:"id"`
}

type PatchUpdateUserMeReq struct {
	g.Meta `path:"/me" tags:"User" method:"patch" summary:"Update User Me"`
	Firstname    string      `json:"firstname"    v:"required"`
	Lastname     string      `json:"lastname"     v:"required"`
	Nationality  string      `json:"nationality"  v:"required"`
}

type PatchUpdateUserMeRes struct {
}

type PatchUpdateUserByIDReq struct {
	g.Meta `path:"/:user_id" tags:"User" method:"patch" summary:"Update User By ID"`
	Firstname    string      `json:"firstname"    v:"required"`
	Lastname     string      `json:"lastname"     v:"required"`
	Nationality  string      `json:"nationality"  v:"required"`
}

type PatchUpdateUserByIDRes struct {
}

type DeleteUserMeReq struct {
	g.Meta `path:"/me" tags:"User" method:"delete" summary:"Delete User Me"`
}

type DeleteUserMeRes struct {
}