// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ScheduleDao is the data access object for table schedule.
type ScheduleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ScheduleColumns // columns contains all the column names of Table for convenient usage.
}

// ScheduleColumns defines and stores column names for table schedule.
type ScheduleColumns struct {
	Id            string //
	ApplicationId string //
	StartTime     string //
	EndTime       string //
	Title         string //
	CompanyId     string //
	UserId        string //
	Location      string //
	Link          string //
	Status        string //
	CreateAt      string // Created Time
	UpdateAt      string // Updated Time
}

// scheduleColumns holds the columns for table schedule.
var scheduleColumns = ScheduleColumns{
	Id:            "id",
	ApplicationId: "application_id",
	StartTime:     "start_time",
	EndTime:       "end_time",
	Title:         "title",
	CompanyId:     "company_id",
	UserId:        "user_id",
	Location:      "location",
	Link:          "link",
	Status:        "status",
	CreateAt:      "create_at",
	UpdateAt:      "update_at",
}

// NewScheduleDao creates and returns a new DAO object for table data access.
func NewScheduleDao() *ScheduleDao {
	return &ScheduleDao{
		group:   "default",
		table:   "schedule",
		columns: scheduleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ScheduleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ScheduleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ScheduleDao) Columns() ScheduleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ScheduleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ScheduleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ScheduleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
