package market

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type (
	sCommodity struct{}
)

func init() {
	service.RegisterCommodity(commodityNew())
}

func commodityNew() *sCommodity {
	return &sCommodity{}
}

func (s *sCommodity) GetCommodityList(ctx context.Context, in model.PageGet) (out []*model.CommodityOut, err error) {
	err = dao.Market.Ctx(ctx).Limit((in.Page-1)*in.RowsPerPage, in.RowsPerPage).Scan(&out)
	if err != nil {
		return
	}
	return
}

func (s *sCommodity) AddCommodity(ctx context.Context, in model.CommodityOut) (err error) {
	/* 	if in.Name == "" {
		return err
	} */

	return dao.Market.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Market.Ctx(ctx).Data(do.Market{
			Id:       in.Id,
			Uid:      in.Uid,
			Name:     in.Name,
			Price:    in.Price,
			Provider: in.Provider,
			Num:      in.Num,
		}).Insert()
		return err
	})
}

func (s *sCommodity) DeleteCommodity(ctx context.Context, uid string) (err error) {
	return dao.Market.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Market.Ctx(ctx).Where(dao.Market.Columns().Uid, uid).Delete()
		return err
	})
}

func (s *sCommodity) ChangeCommodityNumber(ctx context.Context, in model.CommodityChange) (err error) {
	return dao.Market.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Market.Ctx(ctx).Data(do.Market{
			Num: in.Num,
		}).Where(dao.Market.Columns().Uid, in.Uid).Update()
		return err
	})
}
