// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StaffDao is the data access object for table staff.
type StaffDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns StaffColumns // columns contains all the column names of Table for convenient usage.
}

// StaffColumns defines and stores column names for table staff.
type StaffColumns struct {
	Id        string //
	Name      string //
	Gender    string //
	Numbering string //
	Salary    string //
	CardNum   string //
}

// staffColumns holds the columns for table staff.
var staffColumns = StaffColumns{
	Id:        "id",
	Name:      "name",
	Gender:    "gender",
	Numbering: "numbering",
	Salary:    "salary",
	CardNum:   "card_num",
}

// NewStaffDao creates and returns a new DAO object for table data access.
func NewStaffDao() *StaffDao {
	return &StaffDao{
		group:   "default",
		table:   "staff",
		columns: staffColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StaffDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StaffDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StaffDao) Columns() StaffColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StaffDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StaffDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StaffDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
