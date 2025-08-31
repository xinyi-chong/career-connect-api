package template

import (
	"fmt"
	"gf_demo/internal/consts"

	"github.com/gogf/gf/os/genv"
)

func TemplateEmailInvitation(username string, token string) string {
	API_URL := genv.Get(consts.API_URL)
	acceptInvitationLink := fmt.Sprintf("%s/auth/validate?token=%s", API_URL, token)

	return fmt.Sprintf(
		`<mjml>
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
          <mj-text font-weight="bold" font-size="24px" padding="30px 0">
            Validate Your CertService Account
          </mj-text>
          <mj-text>Hello %[1]s,</mj-text>
          <mj-text>
            We would love to have you on board! Please validate your registration with us.
          </mj-text>
					<mj-text>
						<mj-button href="%[2]s" align="center" background-color="#66C7C5">
							Validate
						</mj-button>
					</mj-text>
					<mj-text>
						If you're not able to click on the button above, copy and paste the following link to your browser:<br/> <strong>%[2]s</strong>
					</mj-text>
					<mj-divider border-width="1px" padding-top="24px" />
					<mj-text line-height="24px" padding="24px 10%%">
						If you believe you have received this email in error, please delete it and notify <a href="mailto:%[3]s">%[3]s</a>
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
	  `, username, acceptInvitationLink, genv.Get(consts.SUPPORT_EMAIL),
	)
}
