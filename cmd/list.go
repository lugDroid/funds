package cmd

import (
	"fmt"
	data "go/funds/Data"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the stocks or funds saved",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Accessing saved stocks/funds data...")
		fmt.Println()

		storedFunds := data.ReadFundsFile()

		for _, fund := range storedFunds.Funds {
			fmt.Println("Fund:\t", fund.FundName)
			fmt.Println("Ticker:\t", fund.Ticker)
			fmt.Println("Shares:\t", fund.Shares)
			fmt.Println()
		}
	},
}

func init() {
	RootCommand.AddCommand(listCmd)
}
