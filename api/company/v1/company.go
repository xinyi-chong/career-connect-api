package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type PostCreateCompanyReq struct {
	g.Meta `path:"/" tags:"Company" method:"post" summary:"Create Company"`
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

type PostCreateCompanyRes struct {
	Id					string  		`json:"id"`
}

type GetCompanyMeReq struct {
	g.Meta `path:"/me" tags:"Company" method:"get" summary:"Get Company Me"`
}

type GetCompanyMeRes struct {
	Company *entity.Company `json:"company"`
}

type GetCompanyByIDReq struct {
	g.Meta `path:"/:company_id" tags:"Company" method:"get" summary:"Get Company By ID"`
}

type GetCompanyByIDRes struct {
	Company *entity.Company `json:"company"`
}

type PatchUpdateCompanyMeReq struct {
	g.Meta `path:"/me" tags:"Company" method:"patch" summary:"Update Company Me"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Industry    string      `json:"industry"`
	Tag         string      `json:"tag"`
	Address     string      `json:"address"`
	Website     string      `json:"website"`
	City        string      `json:"city"`
	Size        string      `json:"size"`
	Contact     string      `json:"contact"`
	Logo        *ghttp.UploadFile  		`json:"logo"`
}

type PatchUpdateCompanyMeRes struct {
}

type PatchUpdateCompanyByIDReq struct {
	g.Meta `path:"/:company_id" tags:"Company" method:"patch" summary:"Update Company By ID"`
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

type PatchUpdateCompanyByIDRes struct {}

type DeleteCompanyMeReq struct {
	g.Meta `path:"/me" tags:"Company" method:"delete" summary:"Delete Company Me"`
}

type DeleteCompanyMeRes struct {
}

// CompanyAccounts
type PostCreateCompanyAccountReq struct {
	g.Meta `path:"/account" tags:"Company" method:"post" summary:"Create Company Account"`
	UserID     string      `json:"user_id"     v:"required"`
	RoleID     string      `json:"role_id"     v:"required"`
}

type PostCreateCompanyAccountRes struct {
}

type PostCreateCompanyAccountByCompanyIDReq struct {
	g.Meta `path:"/:company_id/account" tags:"Company" method:"post" summary:"Create Company Account By Company ID"`
	UserID     string      `json:"user_id"     v:"required"`
	RoleID     string      `json:"role_id"     v:"required"`
}

type PostCreateCompanyAccountByCompanyIDRes struct {
	Id string `json:"id"`
}

type GetCompanyAccountsMeReq struct {
	g.Meta `path:"/me/accounts" tags:"Company" method:"get" summary:"Get Company Accounts Me"`
}

type GetCompanyAccountsMeRes struct {
	CompanyAccounts []*entity.CompanyAccounts `json:"company_accounts"`
}

type GetCompanyAccountsByCompanyIDReq struct {
	g.Meta `path:"/:company_id/accounts" tags:"Company" method:"get" summary:"Get Company Accounts By Company ID"`
}

type GetCompanyAccountsByCompanyIDRes struct {
	CompanyAccounts []*entity.CompanyAccounts `json:"company_accounts"`
}

type GetCompanyAccountsByUserIDReq struct {
	g.Meta `path:"/accounts/user/:user_id" tags:"Company" method:"get" summary:"Get Company Accounts By User ID"`
}

type GetCompanyAccountsByUserIDRes struct {
	CompanyAccounts []*entity.CompanyAccounts `json:"company_accounts"`
}

type PatchUpdateCompanyAccountByAccountIDReq struct {
	g.Meta `path:"/account/:account_id" tags:"Company" method:"patch" summary:"Update Company Account By Account ID"`
	RoleID     string      `json:"role_id"     v:"required"`
}

type PatchUpdateCompanyAccountByAccountIDRes struct {
}

type PatchUpdateCompanyAccountByCompanyIDUserIDReq struct {
	g.Meta `path:"/:company_id/user/:user_id/account" tags:"Company" method:"patch" summary:"Update Company Account By Company ID & User ID"`
	RoleID     string      `json:"role_id"     v:"required"`
}

type PatchUpdateCompanyAccountByCompanyIDUserIDRes struct {
}

type DeleteCompanyAccountByAccountIDReq struct {
	g.Meta `path:"/account/:account_id" tags:"Company" method:"delete" summary:"Delete Company Account By Account ID"`
}

type DeleteCompanyAccountByAccountIDRes struct {
}

type DeleteCompanyAccountByCompanyIDUserIDReq struct {
	g.Meta `path:"/:company_id/user/:user_id/account" tags:"Company" method:"delete" summary:"Delete Company Account By Company ID & User ID"`
	RoleID     string      `json:"role_id"     v:"required"`
}

type DeleteCompanyAccountByCompanyIDUserIDRes struct {
}

// Company Subscription
type PostCreateCompanySubscriptionReq struct {
	g.Meta `path:"/subscription" tags:"Company" method:"post" summary:"Create Company Subscription"`
	CompanyID   		string      `json:"company_id"      v:"required"`
	CompanyPlanID   string      `json:"company_plan_id" v:"required"`
	Expiry     		  *gtime.Time `json:"expiry"          v:"required"`
}

type PostCreateCompanySubscriptionRes struct {
	Id     string      `json:"id"`
}

type GetCompanySubscriptionMeReq struct {
	g.Meta `path:"/subscription/me" tags:"Company" method:"get" summary:"Get Company Subscription Me"`
}

type GetCompanySubscriptionMeRes struct {
	CompanySubscription *entity.CompanySubscription `json:"company_subscription"`
}

type PatchUpdateCompanySubscriptionByIDReq struct {
	g.Meta `path:"/subscription/:subscription_id" tags:"Company" method:"patch" summary:"Update Company Subscription By ID"`
	CompanyPlanID   string        `json:"company_plan_id" v:"required"`
	Expiry     		  *gtime.Time   `json:"expiry"          v:"required"`
}

type PatchUpdateCompanySubscriptionByIDRes struct {
}

// Logo
type PostCreateLogoByCompanyIDReq struct {
	g.Meta `path:"/:company_id/logo" tags:"Company" method:"post" summary:"Create Logo By Company ID"`
	Logo        ghttp.UploadFile  		`json:"logo" v:"required"`
}

type PostCreateLogoByCompanyIDRes struct {
}

type PatchUpdateLogoByCompanyIDReq struct {
	g.Meta `path:"/:company_id/logo" tags:"Company" method:"patch" summary:"Update Logo By Company ID"`
	Logo        ghttp.UploadFile  		`json:"logo" v:"required"`
}

type PatchUpdateLogoByCompanyIDRes struct {
}

type DeleteLogoByCompanyIDReq struct {
	g.Meta `path:"/:company_id/logo" tags:"Company" method:"delete" summary:"Delete Logo By Company ID"`
}

type DeleteLogoByCompanyIDRes struct {
}

// Company Plan
type PostCreateCompanyPlanReq struct {
	g.Meta `path:"/plan" tags:"Company" method:"post" summary:"Create Company Plan"`
	Name     string      `json:"name"     v:"required"`
}

type PostCreateCompanyPlanRes struct {
	Id     string      `json:"id"`
}

type GetCompanyPlanMeReq struct {
	g.Meta `path:"/plan/me" tags:"Company" method:"get" summary:"Get Company Plan Me"`
}

type GetCompanyPlanMeRes struct {
	CompanyPlan     *entity.CompanyPlan      `json:"company_plan"`
}

type PatchUpdateCompanyPlanByIDReq struct {
	g.Meta `path:"/plan/:company_plan_id" tags:"Company" method:"patch" summary:"Update Company Plan By ID"`
	Name     string      `json:"name"     v:"required"`
}

type PatchUpdateCompanyPlanByIDRes struct {
}