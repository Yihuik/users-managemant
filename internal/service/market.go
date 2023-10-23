// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
)

type (
	ICommodity interface {
		GetCommodityList(ctx context.Context, in model.PageGet) (out []*model.CommodityOut, err error)
		AddCommodity(ctx context.Context, in model.CommodityOut) (err error)
		DeleteCommodity(ctx context.Context, uid string) (err error)
		ChangeCommodityNumber(ctx context.Context, in model.CommodityChange) (err error)
	}
)

var (
	localCommodity ICommodity
)

func Commodity() ICommodity {
	if localCommodity == nil {
		panic("implement not found for interface ICommodity, forgot register?")
	}
	return localCommodity
}

func RegisterCommodity(i ICommodity) {
	localCommodity = i
}
