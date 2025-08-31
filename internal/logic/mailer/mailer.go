package mailer

import (
	"context"
	"crypto/tls"
	"gf_demo/internal/consts"
	"gf_demo/internal/service"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/Boostport/mjml-go"
	"github.com/go-gomail/gomail"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sMailer struct{}

func init() {
	service.RegisterMailer(New())
}

func New() *sMailer {
	return &sMailer{}
}

func (s *sMailer) SendEmailByMJMLTemplate(ctx context.Context, mjmlTemplate string, subject string, email string) (error) {
	// Send Email
	html, err := mjml.ToHTML(ctx, mjmlTemplate, mjml.WithMinify(true))
	if err != nil {
		g.Log().Error(ctx, consts.FAILED_MJML, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, consts.FAILED_MJML, err.Error())
		return err
	}

	var emptyFile *multipart.FileHeader
	failed, err := s.SendEmail(ctx, html, subject, []string{email}, emptyFile)

	if err != nil {
		return err
	} else if failed > 0 {
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to send email to ", email)
		return err
	}

	return nil
}

func SetUpSMTP(ctx context.Context, from string) (*gomail.Dialer, error) {
	SMTP_FROM := genv.Get(consts.SMTP_FROM)
	SMTP_HOST := genv.Get(consts.SMTP_HOST)
	smtpPortStr := genv.Get(consts.SMTP_PORT)
	SMTP_PASSWORD := genv.Get(consts.SMTP_PASSWORD)
	SMTP_PORT, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to convert SMTP_PORT string to int" + err.Error())
		return nil, err
	}
	
	dialer := gomail.NewDialer(SMTP_HOST, SMTP_PORT, SMTP_FROM, SMTP_PASSWORD)

	// Set up TLS configuration
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_HOST,
	}

	return dialer, nil
}

func (s *sMailer) SendEmail(ctx context.Context, bodyhtml string, subject string, to_emails []string, file *multipart.FileHeader) (int, error) {
	SMTP_FROM := genv.Get(consts.SMTP_FROM)
	dialer, err := SetUpSMTP(ctx, SMTP_FROM)
	if err != nil {
		return 0, err
	}

	failed := 0

	// Set up email message
	for _, to := range to_emails {
		m := gomail.NewMessage()
		m.SetHeader("From", SMTP_FROM)
		m.SetHeader("To", to)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", bodyhtml)

		// Attach the file if it exists
		if file != nil {
			attachment, err := file.Open()
			if err != nil {
				failed++
				err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to open file" + err.Error())
				return failed, err
			}
			defer attachment.Close()

			m.Attach(file.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := io.Copy(w, attachment)
				return err
			}))
		}

		// Send email individually
		if err := dialer.DialAndSend(m); err != nil {
			failed++
			g.Log().Error(ctx, "Failed to send email to ", to, ". Error: ", err)
			continue
		}
	}

	return failed, nil
}
