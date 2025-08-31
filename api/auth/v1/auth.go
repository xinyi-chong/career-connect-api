package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type RegisterUserReq struct {
	g.Meta `path:"/register/user" tags:"Authentication" method:"post" summary:"Register User"`
	Email		     		string   		`json:"email"           v:"required#Please enter email|email#Invalid Email"`
	Password		    string   		`json:"password"        v:"required|length:6,16"`

	Firstname   		string      `json:"firstname"       v:"required#Please enter first name"`
	Lastname 				string      `json:"lastname"        v:"required"`
	Nationality 		string      `json:"nationality"     v:"required"`
	ProfilePicture  *ghttp.UploadFile      `json:"profile_picture"`
}

type RegisterUserRes struct {
	ID        *string  		`json:"id"`
}

type RegisterCompanyReq struct {
	g.Meta `path:"/register/company" tags:"Authentication" method:"post" summary:"Register Company"`
	Email		    string			`json:"email"       v:"required|email"`
	Password		string   		`json:"password"    v:"required|length:6,16"`
	
	Name        string      `json:"name"        v:"required"`
	Description string      `json:"description" v:"required"`
	Industry    string      `json:"industry"    v:"required"`
	Tag         string      `json:"tag"         v:"required"`
	Address     string      `json:"address"     v:"required"`
	Website     string      `json:"website"     v:"required"`
	City        string      `json:"city"        v:"required"`
	Size        string      `json:"size"        v:"required"`
	Contact     string      `json:"contact"     v:"required"`
	Logo        *ghttp.UploadFile  		`json:"logo"`
}

type RegisterCompanyRes struct {
	ID        *string  		`json:"id"`
}

type ActivateAccountReq struct {
	g.Meta `path:"/activate" tags:"Authentication" method:"post" summary:"Activate Account"`
	Email		     		string   		`json:"email"           v:"required#Please enter email|email#Invalid Email"`
}

type ActivateAccountRes struct {
}

type ValidateReq struct {
	g.Meta `path:"/validate" tags:"Authentication" method:"post" summary:"Validate"` // Need Param: token
}

type ValidateRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type SignInCompanyReq struct {
	g.Meta   `path:"/signin/company" method:"post" tags:"Authentication" summary:"Company Sign In 用户登录"`
	Email 	 string	 `json:"email"    v:"required|email"`
	Password string  `json:"password" v:"required"`
}
type SignInCompanyRes struct{
	Company        *entity.Company `json:"company"`
	AccessToken    *string         `json:"access_token"`
	RefreshToken   *string         `json:"refresh_token"`
}

type SignInUserReq struct {
	g.Meta   `path:"/signin/user" method:"post" tags:"Authentication" summary:"User Sign In 用户登录"`
	Email    string `json:"email"    v:"required|email"`
	Password string `json:"password" v:"required"`
}
type SignInUserRes struct{
	User           *entity.User  `json:"user"`
	AccessToken    *string       `json:"access_token"`
	RefreshToken   *string       `json:"refresh_token"`
}

type ForgetPasswordReq struct {
	g.Meta `path:"/forget_password" method:"post" tags:"Authentication" summary:"Forget Password"`
	Email string `json:"email" v:"required|email"`
}
type ForgetPasswordRes struct {}

type ResetPasswordReq struct {
	g.Meta    `path:"/reset_password" method:"post" tags:"Authentication" summary:"Reset Password 重置用户密码"`	// Need Param: token
	Password  string `json:"password"  v:"required|length:6,16"`
	Password2 string `json:"password2" v:"required|length:6,16|same:Password"`
}
type ResetPasswordRes struct {}

type SignOutReq struct {
	g.Meta `path:"/signout" method:"post" tags:"Authentication" summary:"Sign Out 用户登出系统"`
}

type SignOutRes struct{}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh-token" method:"post" tags:"Authentication" summary:"Generate New Access Token By Refresh Token"`
	RefreshToken   string  `json:"refresh_token"  v:"required"`
}

type RefreshTokenRes struct{
	AccessToken   *string  `json:"access_token"`
}