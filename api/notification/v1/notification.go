package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateNotificationReq struct {
	g.Meta   `path:"/" method:"post" tags:"Notification" summary:"Create Notification"`
	RecipientID		string 	 `json:"recipient_id" v:"required"`
	Redirect 			string   `json:"redirect"     v:"required"`
	Title 				string   `json:"title"        v:"required"`
	Description 	string   `json:"description"  v:"required"`
}

type PostCreateNotificationRes struct{
	Id 				*string				`json:"id"`
}

type GetNotificationByIDReq struct {
	g.Meta   `path:"/:notification_id" method:"get" tags:"Notification" summary:"Get Notification By ID"`
}

type GetNotificationByIDRes struct {
	Notification *entity.Notification	`json:"notification"`
}

type GetNotificationByAccountIDReq struct {
	g.Meta   `path:"/account/:account_id" method:"get" tags:"Notification" summary:"Get Notification By Account ID"`
}

type GetNotificationByAccountIDRes struct {
	Notifications []*entity.Notification	`json:"notifications"`
}

type GetNotificationByCompanyIDReq struct {
	g.Meta   `path:"/company/:company_id" method:"get" tags:"Notification" summary:"Get Notification By Company ID"`
}

type GetNotificationByCompanyIDRes struct {
	Notifications []*entity.Notification	`json:"notifications"`
}

type GetNotificationByUserIDReq struct {
	g.Meta   `path:"/user/:user_id" method:"get" tags:"Notification" summary:"Get Notification By User ID"`
}

type GetNotificationByUserIDRes struct {
	Notifications []*entity.Notification	`json:"notifications"`
}

type PatchUpdateNotificationByIDRes struct {}

type PatchUpdateNotificationStatusByIDReq struct { //Update status to seen
	g.Meta   `path:"/:notification_id/status" method:"patch" tags:"Notification" summary:"Update Notification By ID"`
	// Status		string 	 `json:"status" v:"required"`
}

type PatchUpdateNotificationStatusByIDRes struct {}

type DeleteNotificationByIDReq struct {
	g.Meta   `path:"/:notification_id" method:"delete" tags:"Notification" summary:"Delete Notification By ID"`
}

type DeleteNotificationByIDRes struct {}
