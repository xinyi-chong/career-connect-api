package account

import (
	"context"
	v1 "gf_demo/api/account/v1"
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
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type sAccount struct{}

func init() {
	service.RegisterAccount(New())
}

func New() *sAccount {
	return &sAccount{}
}

func (s *sAccount) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account *entity.Account
	err := dao.Account.Ctx(ctx).Where(do.Account{
		Email: email,
	}).Scan(&account)

	if err != nil {
		g.Log().Error(ctx, "Failed to get Account by email", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to get Account By Email: "+err.Error())
		return nil, err
	}

	return account, nil
}

// For Session
func (s *sAccount) GetAccountDetailsByAccountID(ctx context.Context, accountID string) (*model.AccountWithoutPassword, *entity.Company, *entity.User, error) {
	var account *model.AccountWithoutPassword
	err := dao.Account.Ctx(ctx).Where(do.Account{
		Id: accountID,
	}).FieldsEx("password").Scan(&account)
	if err != nil {
		g.Log().Error(ctx, "Failed to Get Account By ID: ", accountID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Account By ID: "+err.Error())
		return nil, nil, nil, err
	}

	company, err := service.Company().GetCompanyByAccountID(ctx, accountID)
	if err != nil {
		return nil, nil, nil, err
	}

	user, err := service.User().GetUserByAccountID(ctx, accountID)
	if err != nil {
		return nil, nil, nil, err
	}

	return account, company, user, nil
}

func (s *sAccount) GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error) {
	var account *entity.Account
	err := dao.Account.Ctx(ctx).Where(do.Account{
		Id: accountID,
	}).Scan(&account)

	if err != nil {
		g.Log().Error(ctx, "Failed to Get Account By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Get Account By ID: "+err.Error())
		return nil, err
	}

	return account, nil
}

func (s *sAccount) CreateAccount(ctx context.Context, email string, password string) (*string, error) {
	account, err := service.Account().GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	} else if account != nil {
		err = gerror.NewCode(gcode.CodeInvalidRequest, "Email already registered.")
		return nil, err
	}

	accountID := uuid.New().String()
	hashedPassword, err := service.Account().HashPassword(ctx, password)
	if err != nil {
		return nil, err
	}

	_, err = dao.Account.Ctx(ctx).Data(do.Account{
		Id:       accountID,
		Email:    email,
		Password: hashedPassword,
		Status:   consts.INACTIVE,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Account: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Account: "+err.Error())
		return nil, err
	}

	g.Log().Info(ctx, consts.SUCCESS_CREATE, "Account: ", accountID)
	return &accountID, nil
}

func (s *sAccount) PatchUpdateEmailMe(ctx context.Context, in model.PatchUpdateEmailMeInput) (newAccessToken *string, err error) {
	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Account.Ctx(ctx).Data(do.Account{
			Email: in.Email,
		}).Where(do.Account{
			Id: in.AccountID,
		}).Update()

		if err != nil {
			g.Log().Error(ctx, "Failed to Update Account Email By ID: ", in.AccountID, err)
			err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Account Email By ID: "+err.Error())
			return err
		}

		g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Account Email for ID: ", in.AccountID)

		// Remove Current Token from Session
		err = service.Session().RemoveCurrentTokenFromSession(ctx)
		if err != nil {
			return err
		}

		// Generate New Access Token to be stored in Session
		newAccessToken, err = service.Token().GenerateAccessToken(in.AccountID, in.Email, in.CompanyID, in.UserID)
		if err != nil {
			return err
		}

		r := g.RequestFromCtx(ctx)

		// Set New Session Data with New Access Token
		r.SetCtxVar(consts.TOKEN, newAccessToken)
		err = service.Session().ResetSessionDataByAccountID(ctx, in.AccountID)
		if err != nil {
			return err
		}

		g.Log().Info(ctx, "Successfully generate new access token for account: ", in.AccountID)
		return nil
	})

	return newAccessToken, err
}

func (s *sAccount) UpdatePasswordByAccountID(ctx context.Context, password string, accountID string) error {
	hashedPassword, err := service.Account().HashPassword(ctx, password)
	if err != nil {
		return err
	}

	_, err = dao.Account.Ctx(ctx).Data(do.Account{
		Password: hashedPassword,
	}).Where(do.Account{
		Id: accountID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Account Password By ID: ", accountID, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Account Password By ID: "+err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, "Account Password for ID: ", accountID)
	return nil
}

func (s *sAccount) PatchUpdatePasswordMe(ctx context.Context, req *v1.PatchUpdatePasswordMeReq, accountID string) error {
	account, err := s.GetAccountByID(ctx, accountID)
	if err != nil {
		return err
	}

	// Verify Current Account Password
	isValid := service.Account().CheckPasswordHash(req.CurrentPassword, account.Password)
	if !isValid {
		g.Log().Error(ctx, "Password is incorrect")
		err = gerror.NewCode(gcode.CodeNotAuthorized, "Password is incorrect")
		return err
	}

	// Update Account Password
	err = s.UpdatePasswordByAccountID(ctx, req.NewPassword, accountID)

	return err
}

func (s *sAccount) DeleteAccountByID(ctx context.Context, accountID string) (err error) {
	_, err = dao.Account.Ctx(ctx).Where(do.Account{
		Id: accountID,
	}).Delete()
	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Account By ID: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Account By ID: "+err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_DELETE, "Account By ID: ", accountID)
	return nil
}

func (s *sAccount) HashPassword(ctx context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		g.Log().Error(ctx, "Failed to Generate Passwored Hash: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Generate Passwored Hash: "+err.Error())
	}

	return string(bytes), err
}

func (s *sAccount) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
