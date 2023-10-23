// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdministratorDao is the data access object for table administrator.
type AdministratorDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdministratorColumns // columns contains all the column names of Table for convenient usage.
}

// AdministratorColumns defines and stores column names for table administrator.
type AdministratorColumns struct {
	Id                    string //
	AdministratorPassport string //
	AdministratorPassword string //
}

// administratorColumns holds the columns for table administrator.
var administratorColumns = AdministratorColumns{
	Id:                    "id",
	AdministratorPassport: "administrator_passport",
	AdministratorPassword: "administrator_password",
}

// NewAdministratorDao creates and returns a new DAO object for table data access.
func NewAdministratorDao() *AdministratorDao {
	return &AdministratorDao{
		group:   "default",
		table:   "administrator",
		columns: administratorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdministratorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdministratorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdministratorDao) Columns() AdministratorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdministratorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdministratorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdministratorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
