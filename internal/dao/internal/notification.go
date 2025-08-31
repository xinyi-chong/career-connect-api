// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationDao is the data access object for table notification.
type NotificationDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NotificationColumns // columns contains all the column names of Table for convenient usage.
}

// NotificationColumns defines and stores column names for table notification.
type NotificationColumns struct {
	Id          string //
	RecipientId string //
	Redirect    string //
	Title       string //
	Description string //
	Status      string //
	CreateAt    string // Created Time
	UpdateAt    string // Updated Time
}

// notificationColumns holds the columns for table notification.
var notificationColumns = NotificationColumns{
	Id:          "id",
	RecipientId: "recipient_id",
	Redirect:    "redirect",
	Title:       "title",
	Description: "description",
	Status:      "status",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewNotificationDao creates and returns a new DAO object for table data access.
func NewNotificationDao() *NotificationDao {
	return &NotificationDao{
		group:   "default",
		table:   "notification",
		columns: notificationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotificationDao) Columns() NotificationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotificationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
