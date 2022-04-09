package cmd

import (
	"fmt"
	data "go/funds/Data"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ticker of stock or fund to delete]",
	Short: "Delete the provided stock or fund",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting fund '%s'\n", args[0])
		fmt.Println()

		funds := data.ReadFundsFile().Funds

		for i, fund := range funds {
			if fund.Ticker == args[0] {
				funds = append(funds[:i], funds[i+1:]...)
				break
			}
		}

		data.WriteFundsFile(data.FundList{
			Funds: funds,
		})

		fmt.Println("Fund/stock deleted successfully")
	},
}

func init() {
	RootCommand.AddCommand(deleteCmd)
}
