// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/dao/internal"
)

// internalCheckCodeDao is internal type for wrapping internal DAO implements.
type internalCheckCodeDao = *internal.CheckCodeDao

// checkCodeDao is the data access object for table check_code.
// You can define custom methods on it to extend its functionality as you wish.
type checkCodeDao struct {
	internalCheckCodeDao
}

var (
	// CheckCode is globally public accessible object for table check_code operations.
	CheckCode = checkCodeDao{
		internal.NewCheckCodeDao(),
	}
)

// Fill with you ideas below.