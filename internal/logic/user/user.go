package user

import (
	"context"
	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) GetUserByAccountID(ctx context.Context, accountID string) (*entity.User, error) {
	var user *entity.User
	err := dao.User.Ctx(ctx).With(
		entity.Media{},
	).Where(do.User{
		AccountId: accountID,
	}).Scan(&user)

	if err != nil {
		g.Log().Error(ctx, "Failed to get User by Account ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to get User by Account ID: " + err.Error())
	}
	
	return user, err
}

func (s *sUser) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	cacheUser := service.Cache().GetCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID, &entity.User{})
	if cacheUser != nil {
		return cacheUser.(*entity.User), nil
	}

	var user *entity.User
	err := dao.User.Ctx(ctx).With(
		entity.Media{},
		entity.Resume{},
		entity.Experience{},
		entity.Certificate{},
		entity.Education{},
		entity.Skill{},
		entity.UserSubscription{},
	).Where(do.User{
		Id: userID,
	}).Scan(&user)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get User By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to get User By ID: " + err.Error())
		return nil, err
	}

	service.Cache().SetCacheWithPrefixByInterface(ctx, consts.CACHE_USER_ID + userID, user)
	
	return user, err
}

func (s *sUser) PostCreateUser(ctx context.Context, req *v1.PostCreateUserReq, accountID string) (id *string, err error) {
	userID := uuid.New().String()
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Media (Profile Picture)
		var mediaID *string
		if req.ProfilePicture != nil {
			if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				mediaID, err = service.Media().CreateMedia(ctx, req.ProfilePicture, accountID)
				return err
			}); err != nil {
				return err
			}
		}

		// Insert User
		if err := dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.User.Ctx(ctx).Data(do.User{
				Id: userID,
				AccountId: accountID,
				Firstname: req.Firstname,
				Lastname: req.Lastname,
				Nationality: req.Nationality,
				ProfilePictureId: mediaID,
			}).Insert()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Create User: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create User: " + err.Error())
			return err
		}
		g.Log().Info(ctx, consts.SUCCESS_CREATE, "user")
		return nil
	})

	return &userID, err
}

func (s *sUser) PatchUpdateUserByID(ctx context.Context, req *v1.PatchUpdateUserByIDReq, userID string) (error) {
	_, err := dao.User.Ctx(ctx).Data(do.User{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Nationality: req.Nationality,
	}).Where(do.User{
		Id: userID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update User By User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update User By User ID: " + err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)
	
	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "User", userID)
	return nil
}

func (s *sUser) DeleteUserByID(ctx context.Context, userID string) (error) {
	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	resumes := user.Resumes
	profilePictureID := user.ProfilePictureId

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Delete Account
		if err := dao.Account.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = service.Account().DeleteAccountByID(ctx, user.AccountId)
			return err
		}); err != nil {
			return err
		}

		// Delete Media (Resume & ProfilePicture)
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// Delete ProfilePicture Media
			if err = service.Media().DeleteMediaByID(ctx, profilePictureID); err != nil {
				return err
			}

			// Delete Resumes Media
			for _, resume := range resumes {
				if err = service.Media().DeleteMediaByID(ctx, resume.MediaId); err != nil {
					return err
				}
			}

			// Remove User Session
			err = service.Session().RemoveSession(ctx)
			return err
		}); err != nil {
			return err
		}

		// Remove Caches
		cacheKeys := []string{
			consts.CACHE_USER_ID + userID,
			consts.CACHE_APPLICATIONS_BY_USER_ID + userID,
			consts.CACHE_NOTIFICATIONS_BY_ACCOUNT_ID + user.AccountId,
			consts.CACHE_COMPANY_ACCOUNTS_BY_USER_ID + userID,
		}
		service.Cache().RemoveMulCachesWithPrefix(ctx, cacheKeys)

		g.Log().Info(ctx, consts.SUCCESS_DELETE, "User: ", userID)
		return nil
	})
	
	return err
}