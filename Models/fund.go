package models

type Fund struct {
	tick          string
	provider      string
	name          string
	price         float32
	changePercent float32
	shares        float32
}

// var (
// 	funds []*Fund
// )

// func AddFund(tick string) Fund {
// 	result := yahoofinanceapi.GetFundData(tick)

// 	newFund := Fund{
// 		tick:          result.symbol,
// 		provider:      result.shortName,
// 		name:          result.longName,
// 		price:         result.regularMarketPrice,
// 		changePercent: result.regularMarketChangePercent,
// 		shares:        0.0,
// 	}

// 	funds = append(funds, &newFund)

// 	return newFund
// }
