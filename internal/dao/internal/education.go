// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EducationDao is the data access object for table education.
type EducationDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns EducationColumns // columns contains all the column names of Table for convenient usage.
}

// EducationColumns defines and stores column names for table education.
type EducationColumns struct {
	Id              string //
	UserId          string //
	StartDate       string //
	EndDate         string //
	InstituteId     string //
	InstituteString string //
	Level           string //
	Programme       string //
	Description     string //
	CreateAt        string // Created Time
	UpdateAt        string // Updated Time
}

// educationColumns holds the columns for table education.
var educationColumns = EducationColumns{
	Id:              "id",
	UserId:          "user_id",
	StartDate:       "start_date",
	EndDate:         "end_date",
	InstituteId:     "institute_id",
	InstituteString: "institute_string",
	Level:           "level",
	Programme:       "programme",
	Description:     "description",
	CreateAt:        "create_at",
	UpdateAt:        "update_at",
}

// NewEducationDao creates and returns a new DAO object for table data access.
func NewEducationDao() *EducationDao {
	return &EducationDao{
		group:   "default",
		table:   "education",
		columns: educationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EducationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EducationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EducationDao) Columns() EducationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EducationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EducationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EducationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
