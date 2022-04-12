package models

type Portfolio struct {
	AssetClasses []AssetClass
	Funds        []Fund
}

type Fund struct {
	FundName string
	AssetId  int
	Ticker   string
	Shares   float32
}

type AssetClass struct {
	Id             int
	Name           string
	PercentOfTotal float32
}
