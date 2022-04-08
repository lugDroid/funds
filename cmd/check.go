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
		fmt.Println()

		data := yahoofinanceapi.GetFundData(args[0])

		if len(data.QuoteResponse.Result) > 0 {
			result := data.QuoteResponse.Result[0]
			fmt.Println("Name:\t\t", result.LongName)
			fmt.Println("Price:\t\t", result.RegularMarketPrice, result.Currency)
			fmt.Printf("Change:\t\t %.2f %%\n", data.QuoteResponse.Result[0].RegularMarketChangePercent)
			fmt.Printf("YTD Return:\t %.2f %%\n", data.QuoteResponse.Result[0].YtdReturn)
		} else {
			fmt.Println("Ticker not valid")
		}
	},
}

func init() {
	RootCommand.AddCommand(checkCmd)
}
