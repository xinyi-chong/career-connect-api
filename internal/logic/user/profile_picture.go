package user

import (
	"context"
	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sUser) PostCreateProfilePicture(ctx context.Context, req *v1.PostCreateProfilePictureReq, userID string, accountID string) (id *string, err error) {
	var mediaID *string
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Media
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			mediaID, err = service.Media().CreateMedia(ctx, req.ProfilePicture, accountID)
			return err
		}); err != nil {
			return err
		}

		// Update User
		if err := dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.User.Ctx(ctx).Data(do.User{
				ProfilePictureId: mediaID,
			}).Where(do.User{
				Id: userID,
			}).Update()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Update Profile Picture ID: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Profile Picture ID: " + err.Error())
			return err
		}

		service.Session().ResetSessionDataByAccountID(ctx, accountID)
		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

		g.Log().Info(ctx, consts.SUCCESS_CREATE, "Profile Picture: ", mediaID)
		return nil
	})

	return mediaID, err
}

func (s *sUser) PatchUpdateProfilePicture(ctx context.Context, req *v1.PatchUpdateProfilePictureReq) error {
	sessionData, err := service.Session().GetSessionDataFromCtx(ctx)
	if err != nil {
		return err
	}

	user := sessionData.User
	err = service.Media().UpdateMediaByID(ctx, user.ProfilePictureId, req.ProfilePicture, sessionData.Account.Id)
	if err != nil {
		return err
	}

	service.Session().ResetSessionDataByAccountID(ctx, sessionData.Account.Id)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + user.Id)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "ProfilePicture: ", user.ProfilePictureId)
	return nil
}

func (s *sUser) DeleteProfilePicture(ctx context.Context) error {
	sessionData, err := service.Session().GetSessionDataFromCtx(ctx)
	if err != nil {
		return err
	}

	err = service.Media().DeleteMediaByID(ctx, sessionData.User.ProfilePictureId)
	if err != nil {
		return err
	}

	service.Session().ResetSessionDataByAccountID(ctx, sessionData.Account.Id)
	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + sessionData.User.Id)

	return err
}
