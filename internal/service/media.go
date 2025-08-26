// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMedia interface {
		CreateMedia(ctx context.Context, file *ghttp.UploadFile, accountID string) (*string, error)
		UpdateMediaByID(ctx context.Context, id string, file *ghttp.UploadFile, accountID string) error
		DeleteMediaByID(ctx context.Context, id string) error
		UploadMedia(ctx context.Context, file ghttp.UploadFile, accountID string) (url *string, objectKey *string, err error)
	}
)

var (
	localMedia IMedia
)

func Media() IMedia {
	if localMedia == nil {
		panic("implement not found for interface IMedia, forgot register?")
	}
	return localMedia
}

func RegisterMedia(i IMedia) {
	localMedia = i
}
