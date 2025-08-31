package template

import (
	"fmt"
	"gf_demo/internal/consts"

	"github.com/gogf/gf/os/genv"
)

func TemplateOTP(username string, otp string, email string) string {

	return fmt.Sprintf(
		`
	<mjml>
		<mj-head>
			<mj-attributes>
				<mj-text line-height="20px" align="center"/>
			</mj-attributes>
		</mj-head>
  
		<mj-body background-color="#1454E2">
			<mj-section>
				<mj-column>
					<mj-spacer height="20px">
					</mj-spacer>
				</mj-column>
			</mj-section>
      <mj-section background-color="#FFFFFF" padding="24px" border-radius="10px">
				<mj-column>
          <mj-text font-weight="bold" font-size="24px" padding="30px">
            You have certificate claim request
          </mj-text>
          <mj-text>Hello %[2]s,</mj-text>
          <mj-text>
            We received your request for claim under email: <strong>%[4]s</strong>. 
            Just a reminder, we'll create your account using this email. 
            Once you log in, this email will serve as your credential. 
            Please make sure to check this email for further instructions.
          </mj-text>
          <mj-text>
            To start this process, you will need to provide this one-time password (OTP) on 
            <strong>CertService</strong>
          </mj-text>
					<mj-text>
						Please enter this number when requested: <strong>%[3]s</strong>
					</mj-text>
					<mj-divider border-width="1px" padding-top="24px" />
					<mj-text line-height="24px" padding="24px 10%%">
						If you believe you have received this email in error, please delete it and notify 
            <a href="mailto:%[5]s">%[5]s</a>
					</mj-text>
				</mj-column>
      </mj-section>
      
      <mj-section>
        <mj-column>
					<mj-text color="#FFFFFF">
						CertService Copyrights © 2024. All right reserved
					</mj-text>
				</mj-column>
			</mj-section>
		</mj-body>
	</mjml>
	`, "config.EMAIL_LOGO_URL", username, otp, email, genv.Get(consts.SUPPORT_EMAIL))
}
