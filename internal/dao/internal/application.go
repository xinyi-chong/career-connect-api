// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplicationDao is the data access object for table application.
type ApplicationDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ApplicationColumns // columns contains all the column names of Table for convenient usage.
}

// ApplicationColumns defines and stores column names for table application.
type ApplicationColumns struct {
	Id         string //
	JobId      string //
	UserId     string //
	Answer     string //
	ApplyAt    string //
	ActivityId string //
	CreateAt   string // Created Time
	UpdateAt   string // Updated Time
}

// applicationColumns holds the columns for table application.
var applicationColumns = ApplicationColumns{
	Id:         "id",
	JobId:      "job_id",
	UserId:     "user_id",
	Answer:     "answer",
	ApplyAt:    "apply_at",
	ActivityId: "activity_id",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}

// NewApplicationDao creates and returns a new DAO object for table data access.
func NewApplicationDao() *ApplicationDao {
	return &ApplicationDao{
		group:   "default",
		table:   "application",
		columns: applicationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplicationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplicationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplicationDao) Columns() ApplicationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplicationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplicationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplicationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
