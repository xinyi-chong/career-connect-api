// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
)

type (
	ISession interface {
		SetSessionDataByToken(ctx context.Context, sessionID string, token string, sessionData model.SessionData) (err error)
		GetSessionDataFromCtx(ctx context.Context) (*model.SessionData, error)
		GetSessionDataByToken(ctx context.Context, token string) (*model.SessionData, error)
		RemoveSession(ctx context.Context) error
		RemoveSessionByID(ctx context.Context, sessionID string) error
		// For Update Email Me
		RemoveCurrentTokenFromSession(ctx context.Context) error
		RemoveCompanySession(ctx context.Context, companyID *string) (err error)
		RemoveCompanyAccountsSessions(ctx context.Context, companyaccounts []*entity.CompanyAccounts) error
		RemoveUserSessionsByRoleID(ctx context.Context, roleID string) error
		ResetSessionDataByAccountID(ctx context.Context, accountID string) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
