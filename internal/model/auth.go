package model

import "github.com/gogf/gf/v2/net/ghttp"

type RegisterCompanyInput struct {
	Email       string
	Password    string
	Name        string
	Description string
	Industry    string
	Tag         string
	Address     string
	Website     string
	City        string
	Size        string
	Contact     string
	Logo        *ghttp.UploadFile
}

type RegisterUserInput struct {
	Email          string
	Password       string
	Firstname      string
	Lastname       string
	Nationality    string
	ProfilePicture *ghttp.UploadFile
}

type SignInInput struct {
	Email    string
	Password string
}

type ResetPasswordInput struct {
	Token    string
	Password string
}
