package cmd

import (
	"fmt"
	yahoofinanceapi "go/funds/YahooFinanceAPI"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [code of stock or fund to check]",
	Short: "Check the value of the provided stock or fund",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving stock/fund information...")
		data := yahoofinanceapi.GetFundData(args[0])

		if len(data.QuoteResponse.Result) > 0 {
			fmt.Println(data.QuoteResponse.Result[0].LongName)
			fmt.Println(data.QuoteResponse.Result[0].RegularMarketPrice)
			fmt.Println(data.QuoteResponse.Result[0].RegularMarketChangePercent)
		} else {
			fmt.Println("Ticker not valid")
		}
	},
}
