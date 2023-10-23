// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InOutDao is the data access object for table in_out.
type InOutDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns InOutColumns // columns contains all the column names of Table for convenient usage.
}

// InOutColumns defines and stores column names for table in_out.
type InOutColumns struct {
	Id        string //
	Uid       string //
	Count     string //
	Time      string //
	Middleman string //
	InPrice   string //
}

// inOutColumns holds the columns for table in_out.
var inOutColumns = InOutColumns{
	Id:        "id",
	Uid:       "uid",
	Count:     "count",
	Time:      "time",
	Middleman: "middleman",
	InPrice:   "in_price",
}

// NewInOutDao creates and returns a new DAO object for table data access.
func NewInOutDao() *InOutDao {
	return &InOutDao{
		group:   "default",
		table:   "in_out",
		columns: inOutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InOutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InOutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InOutDao) Columns() InOutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InOutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InOutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InOutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
