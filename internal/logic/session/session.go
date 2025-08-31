package session

import (
	"context"
	"gf_demo/internal/consts"
	"gf_demo/internal/model"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
	"net/http"

	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

func (s *sSession) SetSessionDataByToken(ctx context.Context, sessionID string, token string, sessionData model.SessionData) (err error) {
	req := g.RequestFromCtx(ctx)

	// id, _ := req.Session.Id()

	// Set gfsessionid to Cookie
	domain := genv.Get(consts.DOMAIN)
	http.SetCookie(req.Response.Writer, &http.Cookie{
    Name:     req.Server.GetSessionIdName(),
    Value:    sessionID,
    Path:     "/",
    Domain:   domain,
	Secure:   true,
	SameSite: http.SameSiteNoneMode,
	})
	// req.Cookie.Set(req.Server.GetSessionIdName(), sessionID)

	// Set SessionData
	req.Session.SetId(sessionID)
	err = req.Session.Set(token, sessionData)
	req.Session.Close()

	if err != nil {
		g.Log().Error(ctx, "Failed to Set Session Data By Token: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Set Session Data By Token: " + err.Error())
	}

	return err
}

func (s *sSession) GetSessionDataFromCtx(ctx context.Context) (*model.SessionData, error) {
	req := g.RequestFromCtx(ctx)
	data := req.GetCtxVar(consts.SESSION_DATA)
	sessionData, ok := data.Val().(*model.SessionData)

	if !ok {
		g.Log().Error(ctx, "Failed to convert Session Data type")
		err := gerror.NewCode(gcode.CodeOperationFailed, "Failed to convert Session Data type")
		return nil, err
	}

	return sessionData, nil
}

func (s *sSession) GetSessionDataByToken(ctx context.Context, token string) (*model.SessionData, error) {
	req := g.RequestFromCtx(ctx)

	// Get Session ID from Cookie
	sessionID := req.Cookie.Get(req.Server.GetSessionIdName()).String()

	req.Session.SetId(sessionID)
	data, err := req.Session.Get(token)
	req.Session.Close()
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Session By Token: ", token, err)
		err := gerror.NewCode(gcode.CodeNotFound, "Failed to Get Session By Token: " + err.Error())
		return nil, err
	}

	sessionData := service.Cache().UnmarshalJson(ctx, data.String(), &model.SessionData{})
	if sessionData == nil {
		g.Log().Error(ctx, "Failed to Unmarshal Session Data: ", token, err)
		err := gerror.NewCode(gcode.CodeOperationFailed, "Failed to Unmarshal Session Data")
		return nil, err
	}
	// sessionData, err := ConvertToSessionDataType(data)
	// if err != nil {
	// 	return nil, err
	// }

	return sessionData.(*model.SessionData), nil
}

// func ConvertToSessionDataType(data *gvar.Var) (*model.SessionData, error) {
// 	interfaceSession := data.Interface()

// 	var sessionData model.SessionData
// 	jsonData, err := json.Marshal(interfaceSession)
// 	if err != nil {
// 		g.Log().Error(context.Background(), "Error marshaling map to JSON: ", err)
// 		return nil, err
// 	}
// 	err = json.Unmarshal(jsonData, &sessionData)
// 	if err != nil {
// 		g.Log().Error(context.Background(), "Error unmarshaling JSON to SessionData type: ", err)
// 		return nil, err
// 	}

// 	return &sessionData, nil
// }

func (s *sSession) RemoveSession(ctx context.Context) error {
	// Get Session ID from Cookies
	req := g.RequestFromCtx(ctx)
	sessionID := req.GetSessionId()

	_, err := g.Redis().Del(ctx, sessionID)
	if err != nil {
		g.Log().Error(ctx, "Failed to Remove Session from Redis: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Remove Session from Redis: " + err.Error())
	}
	
	return err
}

func (s *sSession) RemoveSessionByID(ctx context.Context, sessionID string) error {
	_, err := g.Redis().Del(ctx, sessionID)
	if err != nil {
		g.Log().Error(ctx, "Failed to Remove Session from Redis By Session ID: ", sessionID, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Remove Session from Redis By Session ID: " + err.Error())
	}
	
	return err
}

// For Update Email Me
func (s *sSession) RemoveCurrentTokenFromSession(ctx context.Context) error {
	req := g.RequestFromCtx(ctx)
	token := req.GetCtxVar(consts.TOKEN).String()

	err := req.Session.Remove(token)
	if err != nil {
		g.Log().Error(ctx, "Failed to remove token from session, Token: ", token, err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Remove Token from Session: " + err.Error())
	}
	
	return err
}

func (s *sSession) RemoveCompanySession(ctx context.Context, companyID *string) (err error) {
	tokenData, _ := service.Token().GetTokenDataFromCtxVar(ctx)

	// Reset/Remove Company Session
	if companyID == tokenData.CompanyID {
		// Reset Session Data
		err = service.Session().ResetSessionDataByAccountID(ctx, tokenData.AccountID)
	} else {
		// Remove Account Session
		err = service.Session().RemoveSessionByID(ctx, *companyID)
	}

	return err
}

func (s *sSession) RemoveCompanyAccountsSessions(ctx context.Context, companyaccounts []*entity.CompanyAccounts) (error) {
	for _, companyaccount := range companyaccounts {
		err := s.RemoveSessionByID(ctx, companyaccount.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sSession) RemoveUserSessionsByRoleID(ctx context.Context, roleID string) (error) {
	// Remove User Account Session that are Associated with the Role
	companyaccounts, err := service.Company().GetCompanyAccountsByRoleID(ctx, roleID)
	if err != nil {
		return err
	}
	
	err = s.RemoveCompanyAccountsSessions(ctx, companyaccounts)

	return err
}

func (s *sSession) ResetSessionDataByAccountID(ctx context.Context, accountID string) error {	
	// Retrieve New Session Data
	account, company, user, err := service.Account().GetAccountDetailsByAccountID(ctx, accountID)
	var sessionID string
	if err != nil {
		return err
	}else if account == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Account Not Found")
		return err
	} else if user != nil {
		sessionID = user.Id
	} else {
		sessionID = company.Id
	}

	// Retrieve Current Permissions
	curSessionData, err := service.Session().GetSessionDataFromCtx(ctx)
	if err != nil {
		return err
	}

	// Set New Session Data By Token
	newSessionData := model.SessionData{
		Account: *account,
		User: user,
		Company: company,
		Permissions: curSessionData.Permissions,
	}

	token := g.RequestFromCtx(ctx).GetCtxVar(consts.TOKEN).String()
	err = service.Session().SetSessionDataByToken(ctx, sessionID, token, newSessionData)

	return err
}
