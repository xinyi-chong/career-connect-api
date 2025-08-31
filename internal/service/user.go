// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/model/entity"
)

type (
	IUser interface {
		GetCertificateByID(ctx context.Context, certificateID string) (*entity.Certificate, error)
		GetCertificatesByUserID(ctx context.Context, userID string) ([]*entity.Certificate, error)
		PostCreateCertificate(ctx context.Context, req *v1.PostCreateCertificateReq, userID string) (*string, error)
		PatchUpdateCertificateByID(ctx context.Context, req *v1.PatchUpdateCertificateByIDReq, certificateID string, userID string) error
		DeleteCertificateByID(ctx context.Context, certificateID string, userID string) error
		// CertService
		GetCertificateCertService(ctx context.Context, accountID string) (*map[string]interface{}, error)
		GetEducationsByUserID(ctx context.Context, userID string) ([]*entity.Education, error)
		GetEducationByID(ctx context.Context, educationID string) (*entity.Education, error)
		PostCreateEducation(ctx context.Context, req *v1.PostCreateEducationReq, userID string) (*string, error)
		PatchUpdateEducationByID(ctx context.Context, req *v1.PatchUpdateEducationByIDReq, educationID string, userID string) error
		DeleteEducationByID(ctx context.Context, educationID string, userID string) error
		GetExperiencesByUserID(ctx context.Context, userID string) ([]*entity.Experience, error)
		GetExperienceByID(ctx context.Context, experienceID string) (*entity.Experience, error)
		PostCreateExperience(ctx context.Context, req *v1.PostCreateExperienceReq, userID string) (*string, error)
		PatchUpdateExperienceByID(ctx context.Context, req *v1.PatchUpdateExperienceByIDReq, experienceID string, userID string) error
		DeleteExperienceByID(ctx context.Context, experienceID string, userID string) (err error)
		PostCreateProfilePicture(ctx context.Context, req *v1.PostCreateProfilePictureReq, userID string, accountID string) (id *string, err error)
		PatchUpdateProfilePicture(ctx context.Context, req *v1.PatchUpdateProfilePictureReq) error
		DeleteProfilePicture(ctx context.Context) error
		GetResumesByUserID(ctx context.Context, userID string) ([]*entity.Resume, error)
		GetResumeByID(ctx context.Context, resumeID string) (*entity.Resume, error)
		PostCreateResume(ctx context.Context, req *v1.PostCreateResumeReq, userID string, accountID string) (id *string, err error)
		PatchUpdateResumeByID(ctx context.Context, req *v1.PatchUpdateResumeByIDReq, resumeID string) error
		DeleteResumeByID(ctx context.Context, resumeID string) error
		GetSkillsByUserID(ctx context.Context, userID string) ([]*entity.Skill, error)
		PostCreateSkill(ctx context.Context, req *v1.PostCreateSkillReq, userID string) (*string, error)
		PatchUpdateSkillByID(ctx context.Context, req *v1.PatchUpdateSkillByIDReq, skillID string, userID string) error
		DeleteSkillByID(ctx context.Context, skillID string, userID string) error
		GetUserByAccountID(ctx context.Context, accountID string) (*entity.User, error)
		GetUserByID(ctx context.Context, userID string) (*entity.User, error)
		PostCreateUser(ctx context.Context, req *v1.PostCreateUserReq, accountID string) (id *string, err error)
		PatchUpdateUserByID(ctx context.Context, req *v1.PatchUpdateUserByIDReq, userID string) error
		DeleteUserByID(ctx context.Context, userID string) error
		GetUserPlanByID(ctx context.Context, userPlanID string) (*entity.UserPlan, error)
		PostCreateUserPlan(ctx context.Context, req *v1.PostCreateUserPlanReq) (*string, error)
		PatchUpdateUserPlanByID(ctx context.Context, req *v1.PatchUpdateUserPlanByIDReq, userPlanID string) error
		DeleteUserPlanByID(ctx context.Context, userPlanID string) (err error)
		GetUserSubscriptionByUserID(ctx context.Context, userID string) (*entity.UserSubscription, error)
		GetUserSubscriptionByID(ctx context.Context, userSubscriptionID string) (*entity.UserSubscription, error)
		PostCreateUserSubscription(ctx context.Context, req *v1.PostCreateUserSubscriptionReq, userID string) (*string, error)
		PatchUpdateUserSubscriptionByUserID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionMeReq, userID string) error
		PatchUpdateUserSubscriptionByID(ctx context.Context, req *v1.PatchUpdateUserSubscriptionByIDReq, userSubscriptionID string) error
		DeleteUserSubscriptionByUserID(ctx context.Context, userID string) (err error)
		DeleteUserSubscriptionByID(ctx context.Context, userSubscriptionID string) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
