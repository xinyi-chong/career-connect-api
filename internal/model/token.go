package model

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	AccountID string  `json:"account_id"`
	Email     string  `json:"email"`
	UserID    *string `json:"user_id"`
	CompanyID *string `json:"company_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	AccountID string `json:"account_id"`
	jwt.RegisteredClaims
}

type ValidateAccountToken struct {
	AccountID string `json:"account_id"`
	Type      string `json:"type"` //validate account / reset password
	jwt.RegisteredClaims
}
