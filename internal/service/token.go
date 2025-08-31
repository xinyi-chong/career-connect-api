// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_demo/internal/model"
)

type (
	IToken interface {
		// Generate token
		GenerateAccessAndRefreshToken(ctx context.Context, accountID string) (*string, *string, error)
		GenerateValidateToken(tokenType string, accountID string) (*string, error)
		GenerateAccessToken(accountID string, email string, companyID *string, userID *string) (*string, error)
		GenerateRefreshToken(accountID string) (*string, error)
		ParseValidateAccountToken(validateAccountToken string, tokenType string) (*model.ValidateAccountToken, error)
		ParseAccessToken(accessToken string) (*model.AccessTokenClaims, error)
		ParseRefreshToken(refreshToken string) (*model.RefreshTokenClaims, error)
		GetTokenDataFromCtxVar(ctx context.Context) (accessTokenClaim *model.AccessTokenClaims, err error)
		// Regenerate Access Token By Refresh Token
		RefreshToken(ctx context.Context, refreshToken string) (*string, error)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
