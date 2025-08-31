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

func (s *sUser) GetResumesByUserID(ctx context.Context, userID string) ([]*entity.Resume, error) {
	var resumes []*entity.Resume
	err := dao.Resume.Ctx(ctx).With(
		entity.Media{},
	).Where(do.Resume{
		UserId: userID,
	}).Scan(&resumes)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Resumes: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Resumes: " + err.Error())
	}

	return resumes, err
}

func (s *sUser) GetResumeByID(ctx context.Context, resumeID string) (*entity.Resume, error) {
	var resume *entity.Resume
	err := dao.Resume.Ctx(ctx).With(
		entity.Media{},
	).Where(do.Resume{
		Id: resumeID,
	}).Scan(&resume)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Resume By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Resume By ID: " + err.Error())
	}

	return resume, err
}

func (s *sUser) PostCreateResume(ctx context.Context, req *v1.PostCreateResumeReq, userID string, accountID string) (id *string, err error) {
	resumeID := uuid.New().String()
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Media
		var mediaID *string
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			mediaID, err = service.Media().CreateMedia(ctx, req.Resume, accountID)
			return err
		}); err != nil {
			return err
		}

		// Insert Resume
		if err := dao.Resume.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.Resume.Ctx(ctx).Data(do.Resume{
				Id: resumeID,
				UserId: userID,
				MediaId: mediaID,
			}).Insert()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Create Resume: ", err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Resume: " + err.Error())
			return err
		}

		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + userID)

		g.Log().Info(ctx, consts.SUCCESS_CREATE, "Resume: ", mediaID)
		return nil
	})

	return &resumeID, err
}

func (s *sUser) PatchUpdateResumeByID(ctx context.Context, req *v1.PatchUpdateResumeByIDReq, resumeID string) error {
	resume, err := s.GetResumeByID(ctx, resumeID)
	if err != nil {
		return err
	}

	user, err := service.User().GetUserByID(ctx, resume.UserId)
	if err != nil {
		return err
	}

	err = service.Media().UpdateMediaByID(ctx, resume.MediaId, req.Resume, user.AccountId)
	if err != nil {
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + resume.UserId)

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Resume: ", resumeID)
	return nil
}

func (s *sUser) DeleteResumeByID(ctx context.Context, resumeID string) error {
	resume, err := s.GetResumeByID(ctx, resumeID)
	if err != nil {
		return err
	}

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Delete Resume
		if err := dao.Resume.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.Resume.Ctx(ctx).Where(do.Resume{
				Id: resumeID,
			}).Delete()
			return err
		}); err != nil {
			g.Log().Error(ctx, "Failed to Delete Resume By ID: ", resumeID, err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Resume By ID: " + err.Error())
			return err
		}

		// Delete Media
		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = service.Media().DeleteMediaByID(ctx, resume.MediaId)
			return err
		}); err != nil {
			return err
		}

		service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID + resume.UserId)

		g.Log().Info(ctx, consts.SUCCESS_DELETE, "Resume: ", resumeID)
		return nil
	})

	return err
}