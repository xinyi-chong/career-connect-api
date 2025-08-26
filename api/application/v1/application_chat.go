package v1

import (
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PostCreateApplicationChatByApplicationIDReq struct {
	g.Meta `path:"/:application_id/chat" tags:"Application" method:"post" summary:"Create Application Chat By Application ID"`
	SenderID    	 string      `json:"sender_id"  v:"required"`
	Name     			 string      `json:"name"       v:"required"`
	Message     	 string      `json:"message"    v:"required"`
}

type PostCreateApplicationChatByApplicationIDRes struct {
	Id					*string  		`json:"id"`
}

type GetApplicationChatByApplicationIDReq struct {
	g.Meta `path:"/:application_id/chat" tags:"Application" method:"get" summary:"Get Application Chat By Application ID"`
}

type GetApplicationChatByApplicationIDRes struct {
	ApplicationChats []*entity.ApplicationChatMessage `json:"application_chats"`
}

type DeleteApplicationChatByApplicationIDReq struct {
	g.Meta `path:"/:application_id/chat" tags:"Application" method:"delete" summary:"Delete Application Chat By Application ID"`
}

type DeleteApplicationChatByApplicationIDRes struct {
}