package cmd

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"

	"gf_demo/internal/consts"
	"gf_demo/internal/controller/account"
	"gf_demo/internal/controller/application"
	"gf_demo/internal/controller/auth"
	"gf_demo/internal/controller/company"
	"gf_demo/internal/controller/feature"
	"gf_demo/internal/controller/job"
	"gf_demo/internal/controller/notification"
	"gf_demo/internal/controller/permission"
	"gf_demo/internal/controller/role"
	"gf_demo/internal/controller/suggest_skill"
	"gf_demo/internal/controller/user"
	"gf_demo/internal/service"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MapGCodeToHTTPStatus(code gcode.Code) int {
	switch code.Code() {
	case 51, 53, 54, 60, 66:
		return http.StatusBadRequest
	case 52, 65:
		return http.StatusNotFound
	case 55, 56, 57:
		return http.StatusServiceUnavailable
	case 58, 59:
		return http.StatusNotImplemented
	case 61:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func MiddlewareErrorStatusCode(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	code := gerror.Code(err)

	if code.Code() > 0 { //error code
		statusCode := MapGCodeToHTTPStatus(code)
		r.Response.WriteHeader(statusCode)
	}

	r.Response.WriteJsonExit(g.Map{
		`code`:    code.Code(),
		`message`: err,
		`data`:    r.GetHandlerResponse(),
	})
}

// Validate token
func ValidateJwt(r *ghttp.Request, allowedAccountType *string) (err error) {
	authHeader := r.Header.Get("Authorization")

	var token string
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "Missing Bearer Token")
		return
	}

	// Token is valid if exists in session
	ok, err := r.Session.Contains(token)
	if err != nil {
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Validate Token: "+err.Error())
		return
	} else if !ok {
		err = gerror.NewCode(gcode.CodeNotFound, "Token does not exist in session: "+token)

		return
	}

	accessClaims, err := service.Token().ParseAccessToken(token)
	if err != nil {
		return
	} else if allowedAccountType != nil {
		if *allowedAccountType == consts.USER && accessClaims.UserID == nil {
			err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
			return
		} else if *allowedAccountType == consts.COMPANY && accessClaims.CompanyID == nil {
			err = gerror.NewCode(gcode.CodeNotAuthorized, consts.ERROR_UNAUTHORIZED)
			return
		}
	}

	// Get Session
	sessionData, err := service.Session().GetSessionDataByToken(r.GetCtx(), token)
	if err != nil {
		return
	}

	// Set Context Variable
	r.SetCtxVar(consts.SESSION_DATA, sessionData)
	r.SetCtxVar(consts.TOKEN_CLAIM, accessClaims)
	r.SetCtxVar(consts.TOKEN, token)
	return nil
}

func ValidateAccount(r *ghttp.Request) {
	err := ValidateJwt(r, nil)
	if err != nil {
		r.SetError(err)
		r.Exit()
	}
	r.Middleware.Next()
}

func ValidateUserAccount(r *ghttp.Request) {
	accountType := consts.USER
	err := ValidateJwt(r, &accountType)
	if err != nil {
		r.SetError(err)
		r.Exit()
	}

	r.Middleware.Next()
}

func ValidateCompanyAccount(r *ghttp.Request) {
	accountType := consts.COMPANY
	err := ValidateJwt(r, &accountType)
	if err != nil {
		r.SetError(err)
		r.Exit()
	}
	r.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// Redis Session
			s.SetSessionMaxAge(time.Minute * 15) //same as access token
			s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
			// s.SetSessionIdName("sessionid")		//default: gfsessionid

			// Redis Cache
			redisCache := gcache.NewAdapterRedis(g.Redis())
			g.DB().GetCache().SetAdapter(redisCache)

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(ghttp.MiddlewareCORS)
				group.Middleware(MiddlewareErrorStatusCode)

				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						auth.NewV1().RegisterCompany,
						auth.NewV1().RegisterUser,
						auth.NewV1().ForgetPassword,
						auth.NewV1().ResetPassword,
						auth.NewV1().Validate,
						auth.NewV1().ActivateAccount,
						auth.NewV1().RefreshToken,
						auth.NewV1().SignInUser,
						auth.NewV1().SignInCompany,
					)

					group.Middleware(ValidateAccount) //get token -> accountID
					group.Bind(
						auth.NewV1().SignOut,
					)
				})

				group.Group("/account", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						account.NewV1().PatchUpdatePasswordMe,
						account.NewV1().PatchUpdateEmailMe, //Reset Token & Session Data
					)
				})

				group.Group("/company", func(group *ghttp.RouterGroup) {
					group.Bind(
						company.NewV1().GetCompanyByID,
						company.NewV1().GetCompanyAccountsByCompanyID,
						company.NewV1().GetCompanyAccountsByUserID,
					)

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateAccount)
						group.Bind(
							company.NewV1().PostCreateCompany,
							company.NewV1().PostCreateCompanyAccount,
							company.NewV1().PostCreateCompanyAccountByCompanyID,
							company.NewV1().PostCreateCompanyPlan,
							company.NewV1().PostCreateCompanySubscription,

							company.NewV1().PatchUpdateCompanyAccountByAccountID,
							company.NewV1().PatchUpdateCompanyAccountByCompanyIDUserID,
							company.NewV1().PatchUpdateCompanyByID,
							company.NewV1().PatchUpdateCompanyPlanByID,
							company.NewV1().PatchUpdateCompanySubscriptionByID,
							company.NewV1().DeleteCompanyAccountByAccountID,
							company.NewV1().DeleteCompanyAccountByCompanyIDUserID,

							// Reset/Remove Company Session Data
							company.NewV1().PostCreateLogoByCompanyID,
							company.NewV1().PatchUpdateLogoByCompanyID,
							company.NewV1().DeleteLogoByCompanyID,
						)
					})

					group.Middleware(ValidateCompanyAccount)
					group.Bind(
						company.NewV1().GetCompanyMe,
						company.NewV1().GetCompanyAccountsMe,
						company.NewV1().GetCompanyPlanMe,
						company.NewV1().GetCompanySubscriptionMe,

						// Remove Session
						company.NewV1().PatchUpdateCompanyMe,
						company.NewV1().DeleteCompanyMe,
					)
				})

				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1().GetCertificateByID,
						user.NewV1().GetCertificatesByUserID,
						user.NewV1().GetEducationByID,
						user.NewV1().GetEducationsByUserID,
						user.NewV1().GetExperienceByID,
						user.NewV1().GetExperiencesByUserID,
						user.NewV1().GetResumeByID,
						user.NewV1().GetSkillsByUserID,
						user.NewV1().GetUserByID,
						user.NewV1().GetUserPlanByID,
						user.NewV1().GetUserSubscriptionByID,
					)

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateAccount)
						group.Bind(
							user.NewV1().PostCreateUserPlan,
							user.NewV1().PostCreateUser,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateUserAccount)
						group.Bind(
							user.NewV1().GetCertificateCertService,
							user.NewV1().GetCertificates,
							user.NewV1().GetEducations,
							user.NewV1().GetExperiences,
							user.NewV1().GetResumes,
							user.NewV1().GetSkills,
							user.NewV1().GetUserMe,
							user.NewV1().GetUserSubscriptionMe,

							user.NewV1().PostCreateProfilePicture,
							user.NewV1().PostCreateCertificate,
							user.NewV1().PostCreateEducation,
							user.NewV1().PostCreateExperience,
							user.NewV1().PostCreateResume,
							user.NewV1().PostCreateSkill,
							user.NewV1().PostCreateUserSubscription,

							user.NewV1().PatchUpdateUserMe,
							user.NewV1().PatchUpdateProfilePicture,
							user.NewV1().PatchUpdateCertificateByID,
							user.NewV1().PatchUpdateEducationByID,
							user.NewV1().PatchUpdateExperienceByID,
							user.NewV1().PatchUpdateResumeByID,
							user.NewV1().PatchUpdateSkillByID,
							user.NewV1().PatchUpdateUserPlanByID,
							user.NewV1().PatchUpdateUserSubscriptionByID,
							user.NewV1().PatchUpdateUserSubscriptionMe,

							user.NewV1().DeleteProfilePicture,
							user.NewV1().DeleteCertificateByID,
							user.NewV1().DeleteEducationByID,
							user.NewV1().DeleteExperienceByID,
							user.NewV1().DeleteResumeByID,
							user.NewV1().DeleteSkillByID,
							user.NewV1().DeleteUserPlanByID,
							user.NewV1().DeleteUserSubscriptionByID,
							user.NewV1().DeleteUserSubscriptionMe,
							user.NewV1().DeleteUserMe,
						)
					})
				})

				group.Group("/job", func(group *ghttp.RouterGroup) {
					group.Bind(
						job.NewV1().GetJobsByCompanyID,
						job.NewV1().GetJobsCreatedByCompanyID,
						job.NewV1().GetJobsCreatedByUserID,
						job.NewV1().GetJobQuestionByID,
					)

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateAccount)
						group.Bind(
							job.NewV1().PostCreateJob,
							job.NewV1().PostCreateJobQuestion,

							job.NewV1().PatchUpdateJobByID,
							job.NewV1().PatchUpdateJobQuestionByID,

							job.NewV1().DeleteJobByID,
							job.NewV1().DeleteJobQuestionByID,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateCompanyAccount)
						group.Bind(
							job.NewV1().GetJobByIDCompany,
							job.NewV1().GetJobsMeCompany,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateUserAccount)
						group.Bind(
							job.NewV1().GetJobByIDUser,
							job.NewV1().GetJobsMeUser,
							job.NewV1().GetJobsAllMe,
							job.NewV1().GetJobsAllByKeywords,
						)
					})
				})

				group.Group("/application", func(group *ghttp.RouterGroup) {
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateAccount)
						group.Bind(
							application.NewV1().GetActivityByID,
							application.NewV1().GetApplicationByID,
							application.NewV1().GetApplicationByJobID,
							application.NewV1().GetApplicationChatByApplicationID,
							application.NewV1().GetApplicationFilesByApplicationID,

							application.NewV1().PostCreateActivity,
							application.NewV1().PostCreateActivityByApplicationID,
							application.NewV1().PostCreateApplicationChatByApplicationID,

							application.NewV1().PatchUpdateActivityByID,
							application.NewV1().PatchUpdateApplicationByID,

							application.NewV1().DeleteActivityByID,
							application.NewV1().DeleteApplicationChatByApplicationID,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateUserAccount)
						group.Bind(
							application.NewV1().GetApplicationsMe,
							application.NewV1().PostCreateApplicationByJobID,
							application.NewV1().PostCreateApplicationFilesByApplicationID,
							application.NewV1().PostCreateApplicationFileByApplicationIDResumeID,
							application.NewV1().DeleteApplicationByID,
							application.NewV1().DeleteApplicationFileByApplicationIDFileID,
						)
					})
				})

				group.Group("/schedule", func(group *ghttp.RouterGroup) {
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateAccount)
						group.Bind(
							application.NewV1().GetScheduleByCompanyID,
							application.NewV1().GetScheduleByApplicationID,
							application.NewV1().PostCreateScheduleByApplicationID,
							application.NewV1().PatchUpdateScheduleByID,
							application.NewV1().DeleteScheduleByID,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateCompanyAccount)
						group.Bind(
							application.NewV1().GetScheduleMeCompany,
							application.NewV1().GetScheduleByIDCompany,
							application.NewV1().GetScheduleByApplicationIDCompany,
						)
					})

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(ValidateUserAccount)
						group.Bind(
							application.NewV1().GetScheduleMeUser,
							application.NewV1().GetScheduleByIDUser,
							application.NewV1().GetScheduleByApplicationIDUser,
						)
					})
				})

				group.Group("/suggest_skill", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						suggest_skill.NewV1(),
					)
				})

				group.Group("/notification", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						notification.NewV1(),
					)
				})

				group.Group("/role", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						role.NewV1(),
					)
				})

				group.Group("/feature", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						feature.NewV1(),
					)
				})

				group.Group("/permission", func(group *ghttp.RouterGroup) {
					group.Middleware(ValidateAccount)
					group.Bind(
						permission.NewV1(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
