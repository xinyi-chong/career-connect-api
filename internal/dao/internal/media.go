// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MediaDao is the data access object for table media.
type MediaDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns MediaColumns // columns contains all the column names of Table for convenient usage.
}

// MediaColumns defines and stores column names for table media.
type MediaColumns struct {
	Id       string //
	Url      string //
	Key      string //
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// mediaColumns holds the columns for table media.
var mediaColumns = MediaColumns{
	Id:       "id",
	Url:      "url",
	Key:      "key",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewMediaDao creates and returns a new DAO object for table data access.
func NewMediaDao() *MediaDao {
	return &MediaDao{
		group:   "default",
		table:   "media",
		columns: mediaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MediaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MediaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MediaDao) Columns() MediaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MediaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MediaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MediaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
