// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanySubscriptionDao is the data access object for table company_subscription.
type CompanySubscriptionDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns CompanySubscriptionColumns // columns contains all the column names of Table for convenient usage.
}

// CompanySubscriptionColumns defines and stores column names for table company_subscription.
type CompanySubscriptionColumns struct {
	Id            string //
	CompanyId     string //
	CompanyPlanId string //
	Status        string //
	Expiry        string //
	CreateAt      string // Created Time
	UpdateAt      string // Updated Time
}

// companySubscriptionColumns holds the columns for table company_subscription.
var companySubscriptionColumns = CompanySubscriptionColumns{
	Id:            "id",
	CompanyId:     "company_id",
	CompanyPlanId: "company_plan_id",
	Status:        "status",
	Expiry:        "expiry",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
}

// NewCompanySubscriptionDao creates and returns a new DAO object for table data access.
func NewCompanySubscriptionDao() *CompanySubscriptionDao {
	return &CompanySubscriptionDao{
		group:   "default",
		table:   "company_subscription",
		columns: companySubscriptionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanySubscriptionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanySubscriptionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CompanySubscriptionDao) Columns() CompanySubscriptionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanySubscriptionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanySubscriptionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanySubscriptionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
