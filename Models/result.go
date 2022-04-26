package models

import yahoofinanceapi "go/funds/YahooFinanceAPI"

type result struct {
	Name           string
	Value          float32
	Category       string
	PercentOfTotal float32
}

func CalculateResults(data Portfolio) []result {
	var results []result

	for _, r := range data.Assets {
		assetData := yahoofinanceapi.GetAssetData(r.Ticker)
		assetValue := assetData.QuoteResponse.Result[0].RegularMarketPrice * float64(r.Shares)
		assetCategory := data.Categories[r.Id-1]

		newResult := result{
			Name:           r.Name,
			Value:          float32(assetValue),
			Category:       assetCategory.Name,
			PercentOfTotal: 0,
		}

		results = append(results, newResult)
	}

	totalValue := calculateTotalValue(results)

	for i := range results {
		results[i].PercentOfTotal = (100 * results[i].Value) / totalValue
	}

	return results
}

func calculateTotalValue(results []result) float32 {
	total := float32(0.0)

	for _, r := range results {
		total += r.Value
	}

	return total
}
