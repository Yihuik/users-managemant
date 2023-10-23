// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckCodeDao is the data access object for table check_code.
type CheckCodeDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns CheckCodeColumns // columns contains all the column names of Table for convenient usage.
}

// CheckCodeColumns defines and stores column names for table check_code.
type CheckCodeColumns struct {
	Id          string //
	PhoneNumber string //
	CodeId      string //
	UpdateTime  string //
	EndTime     string //
}

// checkCodeColumns holds the columns for table check_code.
var checkCodeColumns = CheckCodeColumns{
	Id:          "id",
	PhoneNumber: "phone_number",
	CodeId:      "code_id",
	UpdateTime:  "update_time",
	EndTime:     "end_time",
}

// NewCheckCodeDao creates and returns a new DAO object for table data access.
func NewCheckCodeDao() *CheckCodeDao {
	return &CheckCodeDao{
		group:   "default",
		table:   "check_code",
		columns: checkCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CheckCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CheckCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CheckCodeDao) Columns() CheckCodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CheckCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CheckCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CheckCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
