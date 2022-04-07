package yahoofinanceapi

type YFApiResponse struct {
	QuoteResponse struct {
		Result []struct {
			Language                          string  `json:"language"`
			Region                            string  `json:"region"`
			QuoteType                         string  `json:"quoteType"`
			TypeDisp                          string  `json:"typeDisp"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			Triggerable                       bool    `json:"triggerable"`
			CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
			Currency                          string  `json:"currency"`
			YtdReturn                         float64 `json:"ytdReturn"`
			TrailingThreeMonthReturns         float64 `json:"trailingThreeMonthReturns"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
			SourceInterval                    int     `json:"sourceInterval"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			Tradeable                         bool    `json:"tradeable"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Market                            string  `json:"market"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			ShortName                         string  `json:"shortName"`
			LongName                          string  `json:"longName"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			MarketState                       string  `json:"marketState"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			PriceHint                         int     `json:"priceHint"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			Symbol                            string  `json:"symbol"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteResponse"`
}
