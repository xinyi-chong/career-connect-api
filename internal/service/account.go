// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/account/v1"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
)

type (
	IAccount interface {
		GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
		// For Session
		GetAccountDetailsByAccountID(ctx context.Context, accountID string) (*model.AccountWithoutPassword, *entity.Company, *entity.User, error)
		GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error)
		CreateAccount(ctx context.Context, email string, password string) (*string, error)
		PatchUpdateEmailMe(ctx context.Context, in model.PatchUpdateEmailMeInput) (newAccessToken *string, err error)
		UpdatePasswordByAccountID(ctx context.Context, password string, accountID string) error
		PatchUpdatePasswordMe(ctx context.Context, req *v1.PatchUpdatePasswordMeReq, accountID string) error
		DeleteAccountByID(ctx context.Context, accountID string) (err error)
		HashPassword(ctx context.Context, password string) (string, error)
		CheckPasswordHash(password string, hash string) bool
	}
)

var (
	localAccount IAccount
)

func Account() IAccount {
	if localAccount == nil {
		panic("implement not found for interface IAccount, forgot register?")
	}
	return localAccount
}

func RegisterAccount(i IAccount) {
	localAccount = i
}
