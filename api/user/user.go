// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"gf_demo/api/user/v1"
)

type IUserV1 interface {
	GetCertificates(ctx context.Context, req *v1.GetCertificatesReq) (res *v1.GetCertificatesRes, err error)
	GetCertificatesByUserID(ctx context.Context, req *v1.GetCertificatesByUserIDReq) (res *v1.GetCertificatesByUserIDRes, err error)
	GetCertificateByID(ctx context.Context, req *v1.GetCertificateByIDReq) (res *v1.GetCertificateByIDRes, err error)
	PostCreateCertificate(ctx context.Context, req *v1.PostCreateCertificateReq) (res *v1.PostCreateCertificateRes, err error)
	PatchUpdateCertificateByID(ctx context.Context, req *v1.PatchUpdateCertificateByIDReq) (res *v1.PatchUpdateCertificateByIDRes, err error)
	DeleteCertificateByID(ctx context.Context, req *v1.DeleteCertificateByIDReq) (res *v1.DeleteCertificateByIDRes, err error)
	GetCertificateCertService(ctx context.Context, req *v1.GetCertificateCertServiceReq) (res *v1.GetCertificateCertServiceRes, err error)
	GetEducations(ctx context.Context, req *v1.GetEducationsReq) (res *v1.GetEducationsRes, err error)
	GetEducationsByUserID(ctx context.Context, req *v1.GetEducationsByUserIDReq) (res *v1.GetEducationsByUserIDRes, err error)
	GetEducationByID(ctx context.Context, req *v1.GetEducationByIDReq) (res *v1.GetEducationByIDRes, err error)
	PostCreateEducation(ctx context.Context, req *v1.PostCreateEducationReq) (res *v1.PostCreateEducationRes, err error)
	PatchUpdateEducationByID(ctx context.Context, req *v1.PatchUpdateEducationByIDReq) (res *v1.PatchUpdateEducationByIDRes, err error)
	DeleteEducationByID(ctx context.Context, req *v1.DeleteEducationByIDReq) (res *v1.DeleteEducationByIDRes, err error)
	PostCreateExperience(ctx context.Context, req *v1.PostCreateExperienceReq) (res *v1.PostCreateExperienceRes, err error)
	GetExperiences(ctx context.Context, req *v1.GetExperiencesReq) (res *v1.GetExperiencesRes, err error)
	GetExperienceByID(ctx context.Context, req *v1.GetExperienceByIDReq) (res *v1.GetExperienceByIDRes, err error)
	GetExperiencesByUserID(ctx context.Context, req *v1.GetExperiencesByUserIDReq) (res *v1.GetExperiencesByUserIDRes, err error)
	PatchUpdateExperienceByID(ctx context.Context, req *v1.PatchUpdateExperienceByIDReq) (res *v1.PatchUpdateExperienceByIDRes, err error)
	DeleteExperienceByID(ctx context.Context, req *v1.DeleteExperienceByIDReq) (res *v1.DeleteExperienceByIDRes, err error)
	PostCreateProfilePicture(ctx context.Context, req *v1.PostCreateProfilePictureReq) (res *v1.PostCreateProfilePictureRes, err error)
	PatchUpdateProfilePicture(ctx context.Context, req *v1.PatchUpdateProfilePictureReq) (res *v1.PatchUpdateProfilePictureRes, err error)
	DeleteProfilePicture(ctx context.Context, req *v1.DeleteProfilePictureReq) (res *v1.DeleteProfilePictureRes, err error)
	GetResumes(ctx context.Context, req *v1.GetResumesReq) (res *v1.GetResumesRes, err error)
	GetResumeByID(ctx context.Context, req *v1.GetResumeByIDReq) (res *v1.GetResumeByIDRes, err error)
	PostCreateResume(ctx context.Context, req *v1.PostCreateResumeReq) (res *v1.PostCreateResumeRes, err error)
	PatchUpdateResumeByID(ctx context.Context, req *v1.PatchUpdateResumeByIDReq) (res *v1.PatchUpdateResumeByIDRes, err error)
	DeleteResumeByID(ctx context.Context, req *v1.DeleteResumeByIDReq) (res *v1.DeleteResumeByIDRes, err error)
	GetSkills(ctx context.Context, req *v1.GetSkillsReq) (res *v1.GetSkillsRes, err error)
	GetSkillsByUserID(ctx context.Context, req *v1.GetSkillsByUserIDReq) (res *v1.GetSkillsByUserIDRes, err error)
	PostCreateSkill(ctx context.Context, req *v1.PostCreateSkillReq) (res *v1.PostCreateSkillRes, err error)
	PatchUpdateSkillByID(ctx context.Context, req *v1.PatchUpdateSkillByIDReq) (res *v1.PatchUpdateSkillByIDRes, err error)
	DeleteSkillByID(ctx context.Context, req *v1.DeleteSkillByIDReq) (res *v1.DeleteSkillByIDRes, err error)
	GetUserMe(ctx context.Context, req *v1.GetUserMeReq) (res *v1.GetUserMeRes, err error)
	GetUserByID(ctx context.Context, req *v1.GetUserByIDReq) (res *v1.GetUserByIDRes, err error)
	PostCreateUser(ctx context.Context, req *v1.PostCreateUserReq) (res *v1.PostCreateUserRes, err error)
	PatchUpdateUserMe(ctx context.Context, req *v1.PatchUpdateUserMeReq) (res *v1.PatchUpdateUserMeRes, err error)
	PatchUpdateUserByID(ctx context.Context, req *v1.PatchUpdateUserByIDReq) (res *v1.PatchUpdateUserByIDRes, err error)
	DeleteUserMe(ctx context.Context, req *v1.DeleteUserMeReq) (res *v1.DeleteUserMeRes, err error)
	PostCreateUserPlan(ctx context.Context, req *v1.PostCreateUserPlanReq) (res *v1.PostCreateUserPlanRes, err error)
	GetUserPlanByID(ctx context.Context, req *v1.GetUserPlanByIDReq) (res *v1.GetUserPlanByIDRes, err error)
	PatchUpdateUserPlanByID(ctx context.Context, req *v1.PatchUpdateUserPlanByIDReq) (res *v1.PatchUpdateUserPlanByIDRes, err error)
	DeleteUserPlanByID(ctx context.Context, req *v1.DeleteUserPlanByIDReq) (res *v1.DeleteUserPlanByIDRes, err error)
	PostCreateUserSubscription(ctx context.Context, req *v1.PostCreateUserSubscriptionReq) (res *v1.PostCreateUserSubscriptionRes, err error)
	GetUserSubscriptionMe(ctx context.Context, req *v1.GetUserSubscriptionMeReq) (res *v1.GetUserSubscriptionMeRes, err error)
	GetUserSubscriptionByID(ctx context.Context, req *v1.GetUserSubscriptionByIDReq) (res *v1.GetUserSubscriptionByIDRes, err error)
	PatchUpdateUserSubscriptionMe(ctx context.Context, req *v1.PatchUpdateUserSubscriptionMeReq) (res *v1.PatchUpdateUserSubscriptionMeRes, err error)
	PatchUpdateUserSubscriptionByID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionByIDReq) (res *v1.PatchUpdateUserSubscriptionByIDRes, err error)
	DeleteUserSubscriptionMe(ctx context.Context, req *v1.DeleteUserSubscriptionMeReq) (res *v1.DeleteUserSubscriptionMeRes, err error)
	DeleteUserSubscriptionByID(ctx context.Context, req *v1.DeleteUserSubscriptionByIDReq) (res *v1.DeleteUserSubscriptionByIDRes, err error)
}
