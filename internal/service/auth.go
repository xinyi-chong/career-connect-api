// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/auth/v1"
	"gf_demo/internal/model/entity"
)

type (
	IAuth interface {
		RegisterUser(ctx context.Context, in *v1.RegisterUserReq) (id *string, err error)
		// Register Company
		RegisterCompany(ctx context.Context, req *v1.RegisterCompanyReq) (id *string, err error)
		ActivateAccount(ctx context.Context, req *v1.ActivateAccountReq) (err error)
		SignInCompany(ctx context.Context, req *v1.SignInCompanyReq) (*entity.Company, error)
		SignInUser(ctx context.Context, req *v1.SignInUserReq) (*entity.User, error)
		ForgetPassword(ctx context.Context, req *v1.ForgetPasswordReq) (err error)
		Validate(ctx context.Context, accountID string) (err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
