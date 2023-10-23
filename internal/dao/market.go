// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/dao/internal"
)

// internalMarketDao is internal type for wrapping internal DAO implements.
type internalMarketDao = *internal.MarketDao

// marketDao is the data access object for table market.
// You can define custom methods on it to extend its functionality as you wish.
type marketDao struct {
	internalMarketDao
}

var (
	// Market is globally public accessible object for table market operations.
	Market = marketDao{
		internal.NewMarketDao(),
	}
)

// Fill with you ideas below.
