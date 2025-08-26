// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserSubscriptionDao is the data access object for table user_subscription.
type UserSubscriptionDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserSubscriptionColumns // columns contains all the column names of Table for convenient usage.
}

// UserSubscriptionColumns defines and stores column names for table user_subscription.
type UserSubscriptionColumns struct {
	Id         string //
	UserId     string //
	UserPlanId string //
	Status     string //
	Expiry     string //
	CreateAt   string // Created Time
	UpdateAt   string // Updated Time
}

// userSubscriptionColumns holds the columns for table user_subscription.
var userSubscriptionColumns = UserSubscriptionColumns{
	Id:         "id",
	UserId:     "user_id",
	UserPlanId: "user_plan_id",
	Status:     "status",
	Expiry:     "expiry",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}

// NewUserSubscriptionDao creates and returns a new DAO object for table data access.
func NewUserSubscriptionDao() *UserSubscriptionDao {
	return &UserSubscriptionDao{
		group:   "default",
		table:   "user_subscription",
		columns: userSubscriptionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserSubscriptionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserSubscriptionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserSubscriptionDao) Columns() UserSubscriptionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserSubscriptionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserSubscriptionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserSubscriptionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
