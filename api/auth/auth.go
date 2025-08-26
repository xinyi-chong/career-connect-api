// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"gf_demo/api/auth/v1"
)

type IAuthV1 interface {
	RegisterUser(ctx context.Context, req *v1.RegisterUserReq) (res *v1.RegisterUserRes, err error)
	RegisterCompany(ctx context.Context, req *v1.RegisterCompanyReq) (res *v1.RegisterCompanyRes, err error)
	ActivateAccount(ctx context.Context, req *v1.ActivateAccountReq) (res *v1.ActivateAccountRes, err error)
	Validate(ctx context.Context, req *v1.ValidateReq) (res *v1.ValidateRes, err error)
	SignInCompany(ctx context.Context, req *v1.SignInCompanyReq) (res *v1.SignInCompanyRes, err error)
	SignInUser(ctx context.Context, req *v1.SignInUserReq) (res *v1.SignInUserRes, err error)
	ForgetPassword(ctx context.Context, req *v1.ForgetPasswordReq) (res *v1.ForgetPasswordRes, err error)
	ResetPassword(ctx context.Context, req *v1.ResetPasswordReq) (res *v1.ResetPasswordRes, err error)
	SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
}
