package model

type PatchUpdateEmailMeInput struct {
	Email     string
	AccountID string
	CompanyID *string
	UserID    *string
}
