package model

type CommodityOut struct {
	Id       string `json:"id"      description:""`
	Uid      string `json:"uid"     description:""`
	Price    string `json:"price"     description:""`
	Name     string `json:"name"    description:""`
	Num      string `json:"num"     description:""`
	Provider string `json:"provider"     description:""`
}

type CommodityChange struct {
	Uid string `json:"uid"     description:""`
	Num string `json:"num"     description:""`
}
