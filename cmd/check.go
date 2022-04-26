package cmd

import (
	"fmt"
	yahoofinanceapi "go/funds/YahooFinanceAPI"
	"os"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [code of stock or fund to check]",
	Short: "Check the value of the provided stock or fund",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving stock/fund information...")
		fmt.Println()

		assetData, err := yahoofinanceapi.GetAssetData(args[0])
		if err != nil {
			fmt.Println("Ticker code not valid")
			os.Exit(1)
		}

		fmt.Println("Name:\t\t", assetData.LongName)
		fmt.Println("Price:\t\t", assetData.RegularMarketPrice, assetData.Currency)
		fmt.Printf("Change:\t\t %.2f %%\n", assetData.RegularMarketChangePercent)
		fmt.Printf("YTD Return:\t %.2f %%\n", assetData.YtdReturn)
	},
}

func init() {
	RootCommand.AddCommand(checkCmd)
}
