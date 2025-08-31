// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id               string      `json:"id"                 orm:"id"                 ` // User ID
	AccountId        string      `json:"account_id"         orm:"account_id"         ` //
	Firstname        string      `json:"firstname"          orm:"firstname"          ` //
	Lastname         string      `json:"lastname"           orm:"lastname"           ` //
	Nationality      string      `json:"nationality"        orm:"nationality"        ` //
	ProfilePictureId string      `json:"profile_picture_id" orm:"profile_picture_id" ` //
	CreateAt         *gtime.Time `json:"create_at"          orm:"create_at"          ` // Created Time
	UpdateAt         *gtime.Time `json:"update_at"          orm:"update_at"          ` // Updated Time

	ProfilePicture *Media            `json:"profile_picture"   orm:"profile_picture,   with:id=profile_picture_id"` //
	Resumes        []*Resume         `json:"resumes"           orm:"resumes,           with:user_id=id"`            //
	Experiences    []*Experience     `json:"experiences"       orm:"experiences,       with:user_id=id"`            //
	Certificates   []*Certificate    `json:"certificates"      orm:"certificates,      with:user_id=id"`            //
	Educations     []*Education      `json:"educations"        orm:"educations,        with:user_id=id"`            //
	Skills         []*Skill          `json:"skills"            orm:"skills,            with:user_id=id"`            //
	Subscription   *UserSubscription `json:"user_subscription" orm:"user_subscription, with:user_id=id"`            //

	Companyaccounts []*CompanyAccounts `json:"company_accounts"  orm:"company_accounts,  with:user_id=id"`      //
	Notifications   []*Notification    `json:"notifications"     orm:"notifications,     with:recipient_id=id"` //
}
