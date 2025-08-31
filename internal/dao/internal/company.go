// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanyDao is the data access object for table company.
type CompanyDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CompanyColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyColumns defines and stores column names for table company.
type CompanyColumns struct {
	Id          string // company ID
	AccountId   string //
	Name        string //
	Description string //
	Industry    string //
	Tag         string //
	Address     string //
	Website     string //
	City        string //
	Size        string //
	Contact     string //
	LogoId      string //
	CreateAt    string // Created Time
	UpdateAt    string // Updated Time
}

// companyColumns holds the columns for table company.
var companyColumns = CompanyColumns{
	Id:          "id",
	AccountId:   "account_id",
	Name:        "name",
	Description: "description",
	Industry:    "industry",
	Tag:         "tag",
	Address:     "address",
	Website:     "website",
	City:        "city",
	Size:        "size",
	Contact:     "contact",
	LogoId:      "logo_id",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewCompanyDao creates and returns a new DAO object for table data access.
func NewCompanyDao() *CompanyDao {
	return &CompanyDao{
		group:   "default",
		table:   "company",
		columns: companyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CompanyDao) Columns() CompanyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
