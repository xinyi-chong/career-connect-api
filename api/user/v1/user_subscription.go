package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User Subscription
type PostCreateUserSubscriptionReq struct {
	g.Meta `path:"/subscription" tags:"User" method:"post" summary:"Create User Subscription"`
	UserPlanID   string       `json:"user_plan_id"   v:"required"`
	Expiry     	 *gtime.Time  `json:"expiry"         v:"required"`
}

type PostCreateUserSubscriptionRes struct {
	Id     string      `json:"id"`
}

type GetUserSubscriptionMeReq struct {
	g.Meta `path:"/subscription/me" tags:"User" method:"get" summary:"Get User Subscription Me"`
}

type GetUserSubscriptionMeRes struct {
	UserSubscription *entity.UserSubscription `json:"user_subscription"`
}

type GetUserSubscriptionByIDReq struct {
	g.Meta `path:"/subscription/:subscription_id" tags:"User" method:"get" summary:"Get User Subscription By ID"`
}

type GetUserSubscriptionByIDRes struct {
	UserSubscription *entity.UserSubscription `json:"user_subscription"`
}

type PatchUpdateUserSubscriptionMeReq struct {
	g.Meta `path:"/me/subscription" tags:"User" method:"patch" summary:"Update User Subscription Me"`
	UserPlanID   *string        `json:"user_plan_id"`
	Status			 *string        `json:"status"`
	Expiry     	 *gtime.Time    `json:"expiry"`
}

type PatchUpdateUserSubscriptionMeRes struct {
}

type PatchUpdateUserSubscriptionByIDReq struct {
	g.Meta `path:"/subscription/:subscription_id" tags:"User" method:"patch" summary:"Update User Subscription By ID"`
	UserPlanID   *string        `json:"user_plan_id"`
	Status			 *string        `json:"status"`
	Expiry     	 *gtime.Time    `json:"expiry"`
}

type PatchUpdateUserSubscriptionByIDRes struct {
}

type DeleteUserSubscriptionMeReq struct {
	g.Meta `path:"/me/subscription" tags:"User" method:"delete" summary:"Delete User Subscription Me"`
}

type DeleteUserSubscriptionMeRes struct {
}

type DeleteUserSubscriptionByIDReq struct {
	g.Meta `path:"/subscription/:subscription_id" tags:"User" method:"delete" summary:"Delete User Subscription By ID"`
}

type DeleteUserSubscriptionByIDRes struct {
}