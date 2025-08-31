package application

import (
	"context"
	v1 "gf_demo/api/application/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

func (s *sApplication) GetApplicationChatByApplicationID(ctx context.Context, applicationID string) ([]*entity.ApplicationChatMessage, error) {
	var applicationChats []*entity.ApplicationChatMessage
	err := dao.Applicationchatmessage.Ctx(ctx).Where(do.ApplicationChatMessage{
		ApplicationId: applicationID,
	}).Scan(&applicationChats)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Application Chat Message By Application ID: ", applicationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Application Chat Message By Application ID: "+err.Error())
		return nil, err
	}

	return applicationChats, nil
}

func (s *sApplication) PostCreateApplicationChatByApplicationID(ctx context.Context, req *v1.PostCreateApplicationChatByApplicationIDReq, applicationID string) (*string, error) {
	id := uuid.New().String()
	_, err := dao.Applicationchatmessage.Ctx(ctx).Data(do.ApplicationChatMessage{
		Id:            id,
		SenderId:      req.SenderID,
		Name:          req.Name,
		Message:       req.Message,
		ApplicationId: applicationID,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Application Chat Message By Application ID: ", applicationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Application Chat Message By Application ID: "+err.Error())
		return nil, err
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Application Chat Message By Application ID", applicationID)
	return &id, nil
}

func (s *sApplication) DeleteApplicationChatByApplicationID(ctx context.Context, applicationID string) error {
	_, err := dao.Applicationchatmessage.Ctx(ctx).Where(do.ApplicationChatMessage{
		ApplicationId: applicationID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Application Chat Message By Application ID: ", applicationID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Application Chat Message By Application ID: "+err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Application Chat Message By Application ID", applicationID)
	return nil
}
