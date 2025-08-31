package application

import (
	"context"
	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

func (s *sApplication) GetApplicationFileByApplicationIDFileID(ctx context.Context, applicationID string, fileID string) (*entity.ApplicationFile, error) {
	var applicationFile *entity.ApplicationFile
	err := dao.ApplicationFile.Ctx(ctx).With(
		entity.Media{},
		entity.Application{},
	).Where(do.ApplicationFile{
		Id:            fileID,
		ApplicationId: applicationID,
	}).Scan(&applicationFile)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Application File By Application ID & File ID: ", applicationID, fileID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Application File By Application ID & File ID: "+err.Error())
		return nil, err
	}

	return applicationFile, nil
}

func (s *sApplication) GetApplicationFilesByApplicationID(ctx context.Context, applicationID string) ([]*entity.ApplicationFile, error) {
	var applicationFiles []*entity.ApplicationFile
	err := dao.ApplicationFile.Ctx(ctx).With(
		entity.Media{},
	).Where(do.ApplicationFile{
		ApplicationId: applicationID,
	}).Scan(&applicationFiles)
	if err != nil {
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Application Files By Application ID: "+err.Error())
		g.Log().Error(ctx, "Failed to Get Application Files By Application ID: ", applicationID, err)
		return nil, err
	}

	return applicationFiles, nil
}

func CreateApplicationFile(ctx context.Context, in model.CreateApplicationFileInput) (*string, error) {
	id := uuid.New().String()
	_, err := dao.ApplicationFile.Ctx(ctx).Data(do.ApplicationFile{
		Id:            id,
		ApplicationId: in.ApplicationID,
		MediaId:       in.MediaID,
		FileType:      in.FileType,
	}).Insert()
	if err != nil {
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Application File: "+err.Error())
		g.Log().Error(ctx, "Failed to Create Application File: ", err)
	}

	return &id, err
}

func (s *sApplication) PostCreateApplicationFileByApplicationIDResumeID(ctx context.Context, applicationID string, resumeID string) (*string, error) {
	resume, err := service.User().GetResumeByID(ctx, resumeID)
	if err != nil {
		return nil, err
	}

	applicationFile := model.CreateApplicationFileInput{
		ApplicationID: applicationID,
		MediaID:       resume.MediaId,
		FileType:      consts.RESUME,
	}
	id, err := CreateApplicationFile(ctx, applicationFile)
	if err != nil {
		return nil, err
	}

	application, _ := s.GetApplicationByID(ctx, applicationID)
	if application != nil {
		RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)
	}

	return id, nil
}

func (s *sApplication) PostCreateApplicationFilesByApplicationID(ctx context.Context, req *v1.PostCreateApplicationFilesByApplicationIDReq, applicationID string, accountID string) (int, int) {
	// Insert Media
	noOfSuccessResumes := s.UploadMultileApplicationFiles(ctx, req.Resumes, consts.RESUME, accountID, applicationID)
	noOfSuccessOtherFiles := s.UploadMultileApplicationFiles(ctx, req.OtherFiles, consts.OTHER, accountID, applicationID)

	// Remove Caches
	application, _ := s.GetApplicationByID(ctx, applicationID)
	if application != nil {
		RemoveApplicationCache(ctx, &applicationID, application.UserId, application.JobId)
	}

	return noOfSuccessResumes, noOfSuccessOtherFiles
}

func (s *sApplication) UploadMultileApplicationFiles(ctx context.Context, files []*ghttp.UploadFile, fileType string, accountID string, applicationID string) int {
	successUploaded := 0
	for _, file := range files {
		_ = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
			if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				mediaID, err := service.Media().CreateMedia(ctx, file, accountID)
				if err != nil {
					return err
				}

				if err := dao.ApplicationFile.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
					applicationFile := model.CreateApplicationFileInput{
						ApplicationID: applicationID,
						MediaID:       *mediaID,
						FileType:      fileType,
					}
					_, err = CreateApplicationFile(ctx, applicationFile)
					return err
				}); err != nil {
					return err
				}
				successUploaded++
				g.Log().Info(ctx, consts.SUCCESS_CREATE, "Application File, Application ID: ", applicationID, ", Media ID:", mediaID)
				return nil
			}); err != nil {
				return err
			}
			return nil
		})
	}
	g.Log().Info(ctx, "Success Uploaded ", successUploaded, "/", len(files), "files")
	return successUploaded
}

func DeleteApplicationFile(ctx context.Context, applicationID string, fileID string) error {
	_, err := dao.ApplicationFile.Ctx(ctx).Where(do.ApplicationFile{
		Id:            fileID,
		ApplicationId: applicationID,
	}).Delete()
	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Application File: ", applicationID, fileID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Application File: "+err.Error())
	}
	return err
}

func (s *sApplication) DeleteApplicationFileByApplicationIDFileID(ctx context.Context, applicationID string, fileID string) error {
	applicationFile, err := s.GetApplicationFileByApplicationIDFileID(ctx, applicationID, fileID)
	if err != nil {
		return err
	}

	// If the Application File is a User Account's Resume
	// Delete Application File Only
	resumes, err := service.User().GetResumesByUserID(ctx, applicationFile.Application.UserId)
	for _, resume := range resumes {
		if resume.MediaId == applicationFile.MediaId {
			err = DeleteApplicationFile(ctx, applicationID, fileID)
			// Remove Caches
			RemoveApplicationCache(ctx, &applicationID, applicationFile.Application.UserId, applicationFile.Application.JobId)
			g.Log().Info(ctx, consts.SUCCESS_DELETE, "Resume by Application ID: ", applicationID)
			return nil
		}
	}

	// Delete Application File and Media
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Delete Application File
		if err := dao.ApplicationFile.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			err = DeleteApplicationFile(ctx, applicationID, fileID)
			return err
		}); err != nil {
			return err
		}

		if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// Delete Media
			err = service.Media().DeleteMediaByID(ctx, applicationFile.MediaId)
			return err
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	// Remove Caches
	RemoveApplicationCache(ctx, &applicationID, applicationFile.Application.UserId, applicationFile.Application.JobId)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Resume by Application ID: ", applicationID)
	return nil
}
