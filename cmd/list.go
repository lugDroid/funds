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
			var categoryName string
			for _, c := range storedFunds.AssetsAllocation {
				if fund.AssetId == c.Id {
					categoryName = c.Name
					break
				}

				categoryName = "Not Assigned"
			}

			fmt.Println("Fund:\t\t", fund.FundName)
			fmt.Println("Ticker:\t\t", fund.Ticker)
			fmt.Println("Category:\t", categoryName)
			fmt.Println("Shares:\t\t", fund.Shares)
			fmt.Println()
		}
	},
}

func init() {
	RootCommand.AddCommand(listCmd)
}
