// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplicationChatMessageDao is the data access object for table application_chat_message.
type ApplicationChatMessageDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns ApplicationChatMessageColumns // columns contains all the column names of Table for convenient usage.
}

// ApplicationChatMessageColumns defines and stores column names for table application_chat_message.
type ApplicationChatMessageColumns struct {
	Id            string //
	Name          string //
	Message       string //
	SenderId      string //
	ApplicationId string //
	CreateAt      string // Created Time
	UpdateAt      string // Updated Time
}

// applicationChatMessageColumns holds the columns for table application_chat_message.
var applicationChatMessageColumns = ApplicationChatMessageColumns{
	Id:            "id",
	Name:          "name",
	Message:       "message",
	SenderId:      "sender_id",
	ApplicationId: "application_id",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
}

// NewApplicationChatMessageDao creates and returns a new DAO object for table data access.
func NewApplicationChatMessageDao() *ApplicationChatMessageDao {
	return &ApplicationChatMessageDao{
		group:   "default",
		table:   "application_chat_message",
		columns: applicationChatMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplicationChatMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplicationChatMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplicationChatMessageDao) Columns() ApplicationChatMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplicationChatMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplicationChatMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplicationChatMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
