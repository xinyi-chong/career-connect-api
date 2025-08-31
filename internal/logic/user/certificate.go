package user

import (
	"context"
	"encoding/json"
	v1 "gf_demo/api/user/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
	"io"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
)

func (s *sUser) GetCertificateByID(ctx context.Context, certificateID string) (*entity.Certificate, error) {
	var certificate *entity.Certificate
	err := dao.Certificate.Ctx(ctx).Where(do.Certificate{
		Id: certificateID,
	}).Scan(&certificate)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Certificate by ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Certificate by ID"+err.Error())
	}

	return certificate, err
}

func (s *sUser) GetCertificatesByUserID(ctx context.Context, userID string) ([]*entity.Certificate, error) {
	var certificates []*entity.Certificate
	err := dao.Certificate.Ctx(ctx).Where(do.Certificate{
		UserId: userID,
	}).Scan(&certificates)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Certificates by User ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Certificate by User ID"+err.Error())
	}

	return certificates, err
}

func (s *sUser) PostCreateCertificate(ctx context.Context, req *v1.PostCreateCertificateReq, userID string) (*string, error) {
	certificateID := uuid.New().String()
	_, err := dao.Certificate.Ctx(ctx).Data(do.Certificate{
		Id:     certificateID,
		UserId: userID,
		Name:   req.Name,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Certificate", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Certificate"+err.Error())
		return nil, err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID+userID)

	g.Log().Info(consts.SUCCESS_CREATE, "Certificate", certificateID)
	return &certificateID, nil
}

func (s *sUser) PatchUpdateCertificateByID(ctx context.Context, req *v1.PatchUpdateCertificateByIDReq, certificateID string, userID string) error {
	_, err := dao.Certificate.Ctx(ctx).Data(do.Certificate{
		Name: req.Name,
	}).Where(do.Certificate{
		Id:     certificateID,
		UserId: userID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Certificate By ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Certificate"+err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID+userID)

	g.Log().Info(consts.SUCCESS_UPDATE, "Certificate", certificateID)
	return nil
}

func (s *sUser) DeleteCertificateByID(ctx context.Context, certificateID string, userID string) error {
	_, err := dao.Certificate.Ctx(ctx).Where(do.Certificate{
		Id:     certificateID,
		UserId: userID,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Certificate by ID", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Certificate"+err.Error())
		return err
	}

	service.Cache().RemoveCacheWithPrefix(ctx, consts.CACHE_USER_ID+userID)

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Certificate: ", certificateID)
	return nil
}

// CertService
func (s *sUser) GetCertificateCertService(ctx context.Context, accountID string) (*map[string]interface{}, error) {
	url := consts.CERT_SERVICE_API_URL + "/certificate/account/" + accountID

	resp, err := http.Get(url)
	if err != nil {
		g.Log().Error("Failed to send request: %v", err)
		err = gerror.NewCode(gcode.CodeInvalidRequest, "Failed to send request"+err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		g.Log().Error("Failed to Get CertService: ", resp.StatusCode)
		err = gerror.NewCode(gcode.CodeInvalidRequest, "Failed to Get CertService")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Error("Failed to read response body: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to read response body")
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		g.Log().Error("Failed to unmarshal JSON: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to unmarshal JSON: "+err.Error())
		return nil, err
	}

	return &result, err
}
