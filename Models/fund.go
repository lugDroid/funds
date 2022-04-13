package models

type Portfolio struct {
	AssetsAllocation []Asset
	Funds            []Fund
}

type Fund struct {
	FundName string
	AssetId  int
	Ticker   string
	Shares   float32
}

type Asset struct {
	Id        int
	Name      string
	Allocated float32
}
