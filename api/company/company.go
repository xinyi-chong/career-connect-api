// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package company

import (
	"context"

	"gf_demo/api/company/v1"
)

type ICompanyV1 interface {
	PostCreateCompany(ctx context.Context, req *v1.PostCreateCompanyReq) (res *v1.PostCreateCompanyRes, err error)
	GetCompanyMe(ctx context.Context, req *v1.GetCompanyMeReq) (res *v1.GetCompanyMeRes, err error)
	GetCompanyByID(ctx context.Context, req *v1.GetCompanyByIDReq) (res *v1.GetCompanyByIDRes, err error)
	PatchUpdateCompanyMe(ctx context.Context, req *v1.PatchUpdateCompanyMeReq) (res *v1.PatchUpdateCompanyMeRes, err error)
	PatchUpdateCompanyByID(ctx context.Context, req *v1.PatchUpdateCompanyByIDReq) (res *v1.PatchUpdateCompanyByIDRes, err error)
	DeleteCompanyMe(ctx context.Context, req *v1.DeleteCompanyMeReq) (res *v1.DeleteCompanyMeRes, err error)
	PostCreateCompanyAccount(ctx context.Context, req *v1.PostCreateCompanyAccountReq) (res *v1.PostCreateCompanyAccountRes, err error)
	PostCreateCompanyAccountByCompanyID(ctx context.Context, req *v1.PostCreateCompanyAccountByCompanyIDReq) (res *v1.PostCreateCompanyAccountByCompanyIDRes, err error)
	GetCompanyAccountsMe(ctx context.Context, req *v1.GetCompanyAccountsMeReq) (res *v1.GetCompanyAccountsMeRes, err error)
	GetCompanyAccountsByCompanyID(ctx context.Context, req *v1.GetCompanyAccountsByCompanyIDReq) (res *v1.GetCompanyAccountsByCompanyIDRes, err error)
	GetCompanyAccountsByUserID(ctx context.Context, req *v1.GetCompanyAccountsByUserIDReq) (res *v1.GetCompanyAccountsByUserIDRes, err error)
	PatchUpdateCompanyAccountByAccountID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByAccountIDReq) (res *v1.PatchUpdateCompanyAccountByAccountIDRes, err error)
	PatchUpdateCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.PatchUpdateCompanyAccountByCompanyIDUserIDReq) (res *v1.PatchUpdateCompanyAccountByCompanyIDUserIDRes, err error)
	DeleteCompanyAccountByAccountID(ctx context.Context, req *v1.DeleteCompanyAccountByAccountIDReq) (res *v1.DeleteCompanyAccountByAccountIDRes, err error)
	DeleteCompanyAccountByCompanyIDUserID(ctx context.Context, req *v1.DeleteCompanyAccountByCompanyIDUserIDReq) (res *v1.DeleteCompanyAccountByCompanyIDUserIDRes, err error)
	PostCreateCompanySubscription(ctx context.Context, req *v1.PostCreateCompanySubscriptionReq) (res *v1.PostCreateCompanySubscriptionRes, err error)
	GetCompanySubscriptionMe(ctx context.Context, req *v1.GetCompanySubscriptionMeReq) (res *v1.GetCompanySubscriptionMeRes, err error)
	PatchUpdateCompanySubscriptionByID(ctx context.Context, req *v1.PatchUpdateCompanySubscriptionByIDReq) (res *v1.PatchUpdateCompanySubscriptionByIDRes, err error)
	PostCreateLogoByCompanyID(ctx context.Context, req *v1.PostCreateLogoByCompanyIDReq) (res *v1.PostCreateLogoByCompanyIDRes, err error)
	PatchUpdateLogoByCompanyID(ctx context.Context, req *v1.PatchUpdateLogoByCompanyIDReq) (res *v1.PatchUpdateLogoByCompanyIDRes, err error)
	DeleteLogoByCompanyID(ctx context.Context, req *v1.DeleteLogoByCompanyIDReq) (res *v1.DeleteLogoByCompanyIDRes, err error)
	PostCreateCompanyPlan(ctx context.Context, req *v1.PostCreateCompanyPlanReq) (res *v1.PostCreateCompanyPlanRes, err error)
	GetCompanyPlanMe(ctx context.Context, req *v1.GetCompanyPlanMeReq) (res *v1.GetCompanyPlanMeRes, err error)
	PatchUpdateCompanyPlanByID(ctx context.Context, req *v1.PatchUpdateCompanyPlanByIDReq) (res *v1.PatchUpdateCompanyPlanByIDRes, err error)
}
