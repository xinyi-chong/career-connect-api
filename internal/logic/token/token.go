package token

import (
	"context"
	"gf_demo/internal/consts"
	"gf_demo/internal/model"
	"gf_demo/internal/service"
	"time"

	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
)

type sToken struct{}

func init() {
	service.RegisterToken(New())
}

func New() *sToken {
	return &sToken{}
}

// Generate token
func (s *sToken) GenerateAccessAndRefreshToken(ctx context.Context, accountID string) (*string, *string, error) {
	r := g.RequestFromCtx(ctx)

	// Generate JWT Tokens
	refreshToken, err := service.Token().GenerateRefreshToken(accountID)
	if err != nil {
		return nil, nil, err
	}
	accessToken, err := service.Token().RefreshToken(r.GetCtx(), *refreshToken)
	if err != nil {
		return nil, nil, err
	}
	
	return accessToken, refreshToken, nil
}

func (s *sToken) GenerateValidateToken(tokenType string, accountID string) (*string, error) {
	validateClaims := model.ValidateAccountToken{
		AccountID: accountID,
		Type:      tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	signedAccessToken, err := NewToken(validateClaims)
	if err != nil {
		return nil, err
	}

	return signedAccessToken, nil
}

func (s *sToken) GenerateAccessToken(accountID string, email string, companyID *string, userID *string) (*string, error) {
	accessClaims := model.AccessTokenClaims{
		AccountID:   accountID,
		Email:       email,
		// AccountType: accountType,
		UserID: userID,
		CompanyID: companyID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			// IssuedAt:  jwt.NewNumericDate(time.Now()),
			// NotBefore: jwt.NewNumericDate(time.Now()),
			// Issuer:    "test",
			// Subject:   account.Email,
			// ID:        accountID,
			// Audience:  []string{"somebody_else"},
		},
	}

	signedAccessToken, err := NewToken(accessClaims)
	if err != nil {
		return nil, err
	}

	return signedAccessToken, nil
}

func (s *sToken) GenerateRefreshToken(accountID string) (*string, error) {
	refreshClaims := model.RefreshTokenClaims{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	signedRefreshToken, err := NewToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return signedRefreshToken, err
}


func NewToken(claims jwt.Claims) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	signedToken, err := token.SignedString([]byte(genv.Get(consts.SECRET)))
	if err != nil {
		g.Log().Error(context.Background(), "Failed to Sign Token: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Sign Token: " + err.Error())
	}
	
	return &signedToken, err
}

func (s *sToken) ParseValidateAccountToken(validateAccountToken string, tokenType string) (*model.ValidateAccountToken, error) {
	parsedValidateAccountToken, err := jwt.ParseWithClaims(validateAccountToken, &model.ValidateAccountToken{}, func(token *jwt.Token) (interface{}, error) {
	 	return []byte(genv.Get(consts.SECRET)), nil
	})
	if err != nil {
		g.Log().Error(context.Background(), "Invalid Validate Account Token: ", err)
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Validate Account Token: " + err.Error())
		return nil, err
	}
	claim := parsedValidateAccountToken.Claims.(*model.ValidateAccountToken)
	if !parsedValidateAccountToken.Valid || claim.Type != tokenType {
		g.Log().Error(context.Background(), "Invalid Validate Account Token: ", validateAccountToken)
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Validate Account Token")
		return nil, err
	}

	return claim, nil
}
 
func (s *sToken) ParseAccessToken(accessToken string) (*model.AccessTokenClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &model.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
	 	return []byte(genv.Get(consts.SECRET)), nil
	})

	if err != nil {
		g.Log().Error(context.Background(), "Invalid Access Token: ", err)
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Access Token: " + err.Error())
		return nil, err
	}
	
	if !parsedAccessToken.Valid {
		g.Log().Error(context.Background(), "Invalid Access Token: ", accessToken)
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Access Token")
		return nil, err
	}

	return parsedAccessToken.Claims.(*model.AccessTokenClaims), nil
}
 
func (s *sToken) ParseRefreshToken(refreshToken string) (*model.RefreshTokenClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &model.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
	 return []byte(genv.Get(consts.SECRET)), nil
	})
	if err != nil {
		g.Log().Error(context.Background(), "Invalid Refresh Token:  ", err)
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Refresh Token: " + err.Error())
		return nil, err
	}
	if !parsedRefreshToken.Valid {
		g.Log().Error(context.Background(), "Invalid Refresh Token")
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Refresh Token")
		return nil, err
	}

	return parsedRefreshToken.Claims.(*model.RefreshTokenClaims), nil
}

func (s *sToken) GetTokenDataFromCtxVar(ctx context.Context) (accessTokenClaim *model.AccessTokenClaims, err error) {
	claim := g.RequestFromCtx(ctx).GetCtxVar(consts.TOKEN_CLAIM)
	accessTokenClaim, ok := claim.Val().(*model.AccessTokenClaims)

	if !ok {
		g.Log().Error(ctx, "Invalid Access Token")
		err = gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Access Token")
		return
	}
	
	return accessTokenClaim, nil
}

// Regenerate Access Token By Refresh Token
func (s *sToken) RefreshToken(ctx context.Context, refreshToken string) (*string, error) {
	// Validate Refresh Token
	refreshClaims, err := service.Token().ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Validate Account
	permissions := []*model.Permission{}
	account, company, user, err := service.Account().GetAccountDetailsByAccountID(ctx, refreshClaims.AccountID)
	var userID *string
	var companyID *string
	var sessionID string
	if err != nil {
		return nil, err
	}else if account == nil {
		err = gerror.NewCode(gcode.CodeNotFound, "Account Not Found")
		return nil, err
	}else if (user != nil){	
		// User Account
		userID = &user.Id
		sessionID = user.Id
		permissions, err = service.Permission().GetCompanyPermissionsByUserID(ctx, user.Id)
		if(err != nil) {
			return nil, err
		}
	}else {
		companyID = &company.Id
		sessionID = company.Id
	}

	// Generate new access token
	accessToken, err := service.Token().GenerateAccessToken(refreshClaims.AccountID, account.Email, companyID, userID)
	if err != nil {
		return nil, err
	}

	// Set Session Data
	sessionData := model.SessionData{
		Account: *account,
		User: user,
		Company: company,
		Permissions: permissions,
	}

	err = service.Session().SetSessionDataByToken(ctx, sessionID, *accessToken, sessionData)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}