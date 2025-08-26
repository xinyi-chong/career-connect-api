// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/company/v1"
	"gf_demo/internal/model/entity"
)

type (
	ICompany interface {
		GetCompanyByID(ctx context.Context, companyID string) (*entity.Company, error)
		GetCompanyByAccountID(ctx context.Context, accountID string) (*entity.Company, error)
		PostCreateCompany(ctx context.Context, req *v1.PostCreateCompanyReq, accountID string) (companyID *string, err error)
		// Update Company By ID (Without Logo)
		PatchUpdateCompanyByID(ctx context.Context, req *v1.PatchUpdateCompanyByIDReq, companyID string) error
		DeleteCompanyMe(ctx context.Context) (err error)
		GetCompanyAccountsByCompanyID(ctx context.Context, companyID string) ([]*entity.CompanyAccounts, error)
		GetCompanyAccountsByUserID(ctx context.Context, userID string) ([]*entity.CompanyAccounts, error)
		GetCompanyAccountsByRoleID(ctx context.Context, roleID string) ([]*entity.CompanyAccounts, error)
		PostCreateCompanyAccount(ctx context.Context, req *v1.PostCreateCompanyAccountReq, companyID string) error
		PostCreateCompanyAccountByCompanyID(ctx context.Context, req *v1.PostCreateCompanyAccountByCompanyIDReq, companyID string) error
		PatchUpdateCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByCompanyIDUserIDReq, companyID string, userID string) error
		DeleteCompanyAccountByCompanyIDUserID(ctx context.Context, companyID string, userID string) error
		RemoveCompanyAccountCache(ctx context.Context, companyID string, userID string)
		RemoveCompanyAccountsCaches(ctx context.Context, companyaccounts []*entity.CompanyAccounts)
		PostCreateCompanyPlan(ctx context.Context, req *v1.PostCreateCompanyPlanReq) (*string, error)
		PatchUpdateCompanyPlanByID(ctx context.Context, req *v1.PatchUpdateCompanyPlanByIDReq, companyPlanID string) error
		GetCompanySubscriptionByCompanyID(ctx context.Context, companyID string) (*entity.CompanySubscription, error)
		GetCompanySubscriptionByID(ctx context.Context, subscriptionID string) (*entity.CompanySubscription, error)
		PostCreateCompanySubscription(ctx context.Context, req *v1.PostCreateCompanySubscriptionReq, companyID string) (*string, error)
		PatchUpdateCompanySubscriptionByID(ctx context.Context, req *v1.PatchUpdateCompanySubscriptionByIDReq, subscriptionID string) error
		PostCreateLogoByCompanyID(ctx context.Context, req *v1.PostCreateLogoByCompanyIDReq, companyID string) error
		PatchUpdateLogoByCompanyID(ctx context.Context, req *v1.PatchUpdateLogoByCompanyIDReq, companyID string) error
		DeleteLogoByCompanyID(ctx context.Context, companyID string) error
	}
)

var (
	localCompany ICompany
)

func Company() ICompany {
	if localCompany == nil {
		panic("implement not found for interface ICompany, forgot register?")
	}
	return localCompany
}

func RegisterCompany(i ICompany) {
	localCompany = i
}
