// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"mime/multipart"
)

type (
	IMailer interface {
		SendEmailByMJMLTemplate(ctx context.Context, mjmlTemplate string, subject string, email string) error
		SendEmail(ctx context.Context, bodyhtml string, subject string, to_emails []string, file *multipart.FileHeader) (int, error)
	}
)

var (
	localMailer IMailer
)

func Mailer() IMailer {
	if localMailer == nil {
		panic("implement not found for interface IMailer, forgot register?")
	}
	return localMailer
}

func RegisterMailer(i IMailer) {
	localMailer = i
}
