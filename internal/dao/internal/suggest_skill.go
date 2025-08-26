// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SuggestSkillDao is the data access object for table suggest_skill.
type SuggestSkillDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SuggestSkillColumns // columns contains all the column names of Table for convenient usage.
}

// SuggestSkillColumns defines and stores column names for table suggest_skill.
type SuggestSkillColumns struct {
	Id       string //
	Name     string //
	Category string //
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// suggestSkillColumns holds the columns for table suggest_skill.
var suggestSkillColumns = SuggestSkillColumns{
	Id:       "id",
	Name:     "name",
	Category: "category",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewSuggestSkillDao creates and returns a new DAO object for table data access.
func NewSuggestSkillDao() *SuggestSkillDao {
	return &SuggestSkillDao{
		group:   "default",
		table:   "suggest_skill",
		columns: suggestSkillColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SuggestSkillDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SuggestSkillDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SuggestSkillDao) Columns() SuggestSkillColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SuggestSkillDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SuggestSkillDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SuggestSkillDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
