package auth

import (
	"context"
	v1 "gf_demo/api/auth/v1"
	v1Company "gf_demo/api/company/v1"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model"
	"gf_demo/internal/model/do"
	"gf_demo/internal/model/entity"
	"gf_demo/internal/service"
	"gf_demo/internal/template"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

func (s *sAuth) RegisterUser(ctx context.Context, in *v1.RegisterUserReq) (id *string, err error) {
	var accountID *string
	userID := uuid.New().String()

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Account
		if err := dao.Account.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			accountID, err = service.Account().CreateAccount(ctx, in.Email, in.Password)
			return err
		}); err != nil {
			return err
		}

		// Insert Media (Profile Picture)
		var mediaID *string
		if in.ProfilePicture != nil {
			if err := dao.Media.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				mediaID, err = service.Media().CreateMedia(ctx, in.ProfilePicture, *accountID)
				return err
			}); err != nil {
				return err
			}
		}

		// Insert User
		if err := dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			_, err = dao.User.Ctx(ctx).Data(do.User{
				Id: userID,
				AccountId: *accountID,
				Firstname: in.Firstname,
				Lastname: in.Lastname,
				Nationality: in.Nationality,
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

	if err != nil {
		return nil, err
	}

	// Send Email
	err = SendValidateAccountEmail(ctx, consts.VALIDATE_ACCOUNT_TOKEN, in.Email, *accountID, in.Firstname + in.Lastname)
	return &userID, err
}

// Register Company
func (s *sAuth) RegisterCompany(ctx context.Context, req *v1.RegisterCompanyReq) (id *string, err error) {
	var accountID *string
	var companyID *string

	err = gdb.DB.Transaction(g.DB(), ctx, func(ctx context.Context, tx gdb.TX) error {
		// Insert Account
		if err := dao.Account.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			accountID, err = service.Account().CreateAccount(ctx, req.Email, req.Password)
			return err
		}); err != nil {
			return err
		}

		company := &v1Company.PostCreateCompanyReq{
			Name: req.Name,
			Description: req.Description,
			Industry: req.Industry,
			Tag: req.Tag,
			Address: req.Address,
			Website: req.Website,
			City: req.City,
			Size: req.Size,
			Contact: req.Contact,
			Logo: req.Logo,
		}
		
		companyID, err = service.Company().PostCreateCompany(ctx, company, *accountID)
		if err != nil {
			return err
		}
		g.Log().Info(ctx, consts.SUCCESS_CREATE, "Company: ", companyID)
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Send Email
	err = SendValidateAccountEmail(ctx, consts.VALIDATE_ACCOUNT_TOKEN, req.Email, *accountID, req.Name)
	return companyID, err
}

func (s *sAuth) ActivateAccount(ctx context.Context, req *v1.ActivateAccountReq) (err error) {
	account, err := service.Account().GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return err	
	} 

	user, err := service.User().GetUserByAccountID(ctx, account.Id)
	var name string

	if err != nil {
		return err	
	} else if user != nil {
		name = user.Firstname + user.Lastname
	} else {
		company, err := service.Company().GetCompanyByAccountID(ctx, account.Id)
		if err != nil {
			return err
		}
		name = company.Name
	}

	err = SendValidateAccountEmail(ctx, consts.VALIDATE_ACCOUNT_TOKEN, req.Email, account.Id, name)
	return err

}

func (s *sAuth) SignInCompany(ctx context.Context, req *v1.SignInCompanyReq) (*entity.Company, error) {
	accountID, err := SignIn(ctx, model.SignInInput{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// Is a Company Account ?
	company, _ := service.Company().GetCompanyByAccountID(ctx, *accountID)
	if company == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Not a Company Account")
		return nil, err
	}

	return company, nil
}

func (s *sAuth) SignInUser(ctx context.Context, req *v1.SignInUserReq) (*entity.User, error) {
	accountID, err := SignIn(ctx, model.SignInInput{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// Is a User Account ?
	user, _ := service.User().GetUserByAccountID(ctx, *accountID)
	if user == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Not a User Account")
		return nil, err
	}

	return user, nil
}

func (s *sAuth) ForgetPassword(ctx context.Context, req *v1.ForgetPasswordReq) (err error) {
	account, err := service.Account().GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return err
	} else if account == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Your account has not been registered.")
		return err
	}

	// Generate Reset Password Token
	token, err := service.Token().GenerateValidateToken(consts.RESET_PASSWORD_TOKEN, account.Id)
	if err != nil {
		return err
	}

	// Send Email
	err = service.Mailer().SendEmailByMJMLTemplate(ctx, template.TemplateForgotPassword(req.Email, *token), "Reset Password", req.Email)
	
	return err
}

func (s *sAuth) Validate(ctx context.Context, accountID string) (err error) {
	//Update account status to active
	_, err = dao.Account.Ctx(ctx).Data(do.Account{
		Status: consts.ACTIVE,
	}).Where(do.Account{
		Id: accountID,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Account Status to Active: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Account Status to Active: " + err.Error())
	}

	return err
}

func SignIn(ctx context.Context, in model.SignInInput) (*string, error) {
	account, err := service.Account().GetAccountByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	} else if account == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Email not yet registered.")
		return nil, err
	}	else if account.Status == consts.INACTIVE {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "Account is inactive.")
		return nil, err
	}

	isValid := service.Account().CheckPasswordHash(in.Password, account.Password)
	if !isValid {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "Email or Password is incorrect.")
		return nil, err
	}
	
	return &account.Id, nil
}

func SendValidateAccountEmail(ctx context.Context, tokenType string, email string, accountID string, name string) error {
	// Generate Validate Account Token
	token, err := service.Token().GenerateValidateToken(tokenType, accountID)
	if err != nil {
		return err
	}
	
	// Send Email
	err = service.Mailer().SendEmailByMJMLTemplate(ctx, template.TemplateEmailInvitation(name, *token), "Welcome to FirstLinks!", email)

	return err
}