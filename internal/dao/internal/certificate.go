// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CertificateDao is the data access object for table certificate.
type CertificateDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CertificateColumns // columns contains all the column names of Table for convenient usage.
}

// CertificateColumns defines and stores column names for table certificate.
type CertificateColumns struct {
	Id       string // ID
	UserId   string // user ID
	Name     string // user email
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// certificateColumns holds the columns for table certificate.
var certificateColumns = CertificateColumns{
	Id:       "id",
	UserId:   "user_id",
	Name:     "name",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewCertificateDao creates and returns a new DAO object for table data access.
func NewCertificateDao() *CertificateDao {
	return &CertificateDao{
		group:   "default",
		table:   "certificate",
		columns: certificateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CertificateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CertificateDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CertificateDao) Columns() CertificateColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CertificateDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CertificateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CertificateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
