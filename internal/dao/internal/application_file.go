// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplicationFileDao is the data access object for table application_file.
type ApplicationFileDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns ApplicationFileColumns // columns contains all the column names of Table for convenient usage.
}

// ApplicationFileColumns defines and stores column names for table application_file.
type ApplicationFileColumns struct {
	Id            string //
	ApplicationId string //
	MediaId       string //
	FileType      string //
	CreateAt      string // Created Time
	UpdateAt      string // Updated Time
}

// applicationFileColumns holds the columns for table application_file.
var applicationFileColumns = ApplicationFileColumns{
	Id:            "id",
	ApplicationId: "application_id",
	MediaId:       "media_id",
	FileType:      "file_type",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
}

// NewApplicationFileDao creates and returns a new DAO object for table data access.
func NewApplicationFileDao() *ApplicationFileDao {
	return &ApplicationFileDao{
		group:   "default",
		table:   "application_file",
		columns: applicationFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplicationFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplicationFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplicationFileDao) Columns() ApplicationFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplicationFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplicationFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplicationFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
