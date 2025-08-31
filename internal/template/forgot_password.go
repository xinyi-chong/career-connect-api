package template

import (
	"fmt"
	"gf_demo/internal/consts"

	"github.com/gogf/gf/os/genv"
)

func TemplateForgotPassword(username, token string) string {
	API_URL := genv.Get(consts.API_URL)
	forgotPasswordLink := API_URL + "/auth/reset_password?token=" + token

	return fmt.Sprintf(
		`<mjml>
		<mj-body background-color="#1454E2">
			<mj-section>
				<mj-column>
					<mj-spacer height="20px">
					</mj-spacer>
				</mj-column>
			</mj-section>
		  <mj-section background-color="#ffffff" padding="20px" border-radius="10px">
			<mj-column>
			  <mj-text font-size="24px" font-weight="bold">Hi, %[1]s</mj-text>
			  <mj-text color="#000000">You've requested to reset your password. Click the button below to get started.</mj-text>
				<mj-text>
			 		<mj-button background-color="#66C7C5" color="#ffffff" font-size="20px" href="%[2]s">Reset My Password</mj-button>
				</mj-text>
			  <mj-text color="#626262">If you're not able to click on the button above, copy and paste the following link to your browser:</mj-text>
			  <mj-text color="#5e5e5e" font-size="12px">%[2]s</mj-text>
			
				<mj-divider border-width="1px" padding-top="24px" />
				<mj-text line-height="24px" padding="24px 10%%">
					If you believe you have received this email in error, please delete it and notify <a href="mailto:%[3]s">%[3]s</a>
				</mj-text>
			</mj-column>
		  </mj-section>
			<mj-section>
				<mj-column>
					<mj-text color="#FFFFFF" align="center">
						CertService Copyrights © 2024. All right reserved
					</mj-text>
				</mj-column>
			</mj-section>
		</mj-body>
	  </mjml>`, username, forgotPasswordLink, genv.Get(consts.SUPPORT_EMAIL),
	)
}
