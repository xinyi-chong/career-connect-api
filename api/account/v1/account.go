package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PatchUpdateEmailMeReq struct {
	g.Meta `path:"/email/me" tags:"Account" method:"patch" summary:"Update Account Email"`
	Email  string `json:"email"           v:"required#Please enter email|email#Invalid Email"`
}

type PatchUpdateEmailMeRes struct {
	AccessToken *string `json:"access_token"`
}

type PatchUpdatePasswordMeReq struct {
	g.Meta          `path:"/password/me" tags:"Account" method:"patch" summary:"Update Password"`
	CurrentPassword string `json:"current_password" v:"required|length:6,16"`
	NewPassword     string `json:"new_password"     v:"required|length:6,16"`
	NewPassword2    string `json:"new_password2"    v:"required|length:6,16|same:NewPassword"`
}

type PatchUpdatePasswordMeRes struct {
}
