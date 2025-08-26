// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanyPlanDao is the data access object for table company_plan.
type CompanyPlanDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CompanyPlanColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyPlanColumns defines and stores column names for table company_plan.
type CompanyPlanColumns struct {
	Id       string //
	Name     string //
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// companyPlanColumns holds the columns for table company_plan.
var companyPlanColumns = CompanyPlanColumns{
	Id:       "id",
	Name:     "name",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewCompanyPlanDao creates and returns a new DAO object for table data access.
func NewCompanyPlanDao() *CompanyPlanDao {
	return &CompanyPlanDao{
		group:   "default",
		table:   "company_plan",
		columns: companyPlanColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyPlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyPlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CompanyPlanDao) Columns() CompanyPlanColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyPlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyPlanDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyPlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
