package models

type Portfolio struct {
	AssetAllocation []Asset
	Funds           []Fund
}

type Fund struct {
	FundName string
	Class    AssetClass
	Ticker   string
	Shares   float32
}

type Asset struct {
	Class  AssetClass
	Amount float32
}

type AssetClass int

const (
	Stocks = iota
	Bonds
	Cash
	Gold
)
