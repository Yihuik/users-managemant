package market

import (
	"context"
	"fmt"

	v1 "github.com/gogf/gf-demo-user/v2/api/market/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCommodityController struct{}

func CommodityNew() *sCommodityController {
	return &sCommodityController{}
}

func (c *sCommodityController) GetCommodityList(ctx context.Context, req *v1.GetCommodityListReq) (res v1.GetCommodityListRes, err error) {
	out, err := service.Commodity().GetCommodityList(ctx, model.PageGet{
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
	})
	if err != nil {
		return
	}
	fmt.Println(out)

	var commodityList []*v1.CommodityRes

	if err = gconv.Scan(out, &commodityList); err != nil {
		return
	}

	res.CommodityList = commodityList
	return
}

func (c *sCommodityController) AddCommodity(ctx context.Context, req *v1.AddCommodityReq) (res *v1.AddCommodityRes, err error) {
	err = service.Commodity().AddCommodity(ctx, model.CommodityOut{
		Name:     req.Name,
		Uid:      req.Uid,
		Num:      req.Num,
		Provider: req.Provider,
		Price:    req.Price,
		Id:       req.Id,
	})

	return
}

func (c *sCommodityController) DeleteCommodity(ctx context.Context, req *v1.DeleteCommodityReq) (res *v1.DeleteCommodityRes, err error) {
	err = service.Commodity().DeleteCommodity(ctx, req.Uid)
	return
}

func (c *sCommodityController) ChangeCommodityNumber(ctx context.Context, req *v1.ChangeCommodityNumberReq) (res *v1.ChangeCommodityNumberRes, err error) {
	err = service.Commodity().ChangeCommodityNumber(ctx, model.CommodityChange{
		Uid: req.Uid,
		Num: req.Num,
	})
	return
}
