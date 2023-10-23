package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetCommodityListReq struct {
	g.Meta      `path:"/GetCommodityList" method:"get" summary:""`
	Page        int `v:"required"`
	RowsPerPage int `v:"required"`
}

type GetCommodityListRes struct {
	CommodityList []*CommodityRes `json:"commodityList"`
}

type AddCommodityReq struct {
	g.Meta   `path:"/AddCommodity" method:"post" summary:""`
	Id       string `v:"required"`
	Name     string `v:"required"`
	Num      string `v:"required"`
	Uid      string `v:"required"`
	Provider string `v:"required"`
	Price    string `v:"required"`
}

type AddCommodityRes struct {
	GoodsList []*CommodityRes `json:"commodityList"`
}

type DeleteCommodityReq struct {
	g.Meta `path:"/DeleteCommodity" method:"post" summary:""`
	Uid    string `v:"required"`
}

type DeleteCommodityRes struct {
	CommodityList []*CommodityRes `json:"commodityList"`
}

type ChangeCommodityNumberReq struct {
	g.Meta `path:"/ChangeCommodityNumber" method:"post" summary:""`
	Uid    string `v:"required"`
	Num    string `v:"required"`
}

type ChangeCommodityNumberRes struct {
	CommodityList []*CommodityRes `json:"commodityList"`
}

type AccountReq struct {
}

type AccountRes struct {
}
type CommodityRes struct {
	Uid      string `json:"uid"     description:""`
	Name     string `json:"name"   description:""`
	Num      string `json:"num" description:""`
	Id       string `json:"id"     description:""`
	Price    string `json:"price"     description:""`
	Provider string `json:"provider"     description:""`
}
