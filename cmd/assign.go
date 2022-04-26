package cmd

import (
	"fmt"
	data "go/funds/Data"
	"strconv"

	"github.com/spf13/cobra"
)

var assignCmd = &cobra.Command{
	Use:   "assign [ticker of existing stock or fund] [id of the category to be assigned]",
	Short: "Assign the provided category to a existing stock or fund",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Assigning categories...")
		fmt.Println()

		id, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			fmt.Println("Asset Id required (int)")
			return
		}

		portfolio := data.ReadStorageFile()

		for i, asset := range portfolio.Assets {
			if asset.Ticker == args[0] {
				portfolio.Assets[i].Id = int(id)

				data.WriteStorageFile(portfolio)
				fmt.Println("Category successfully assigned")

				return
			}
		}

		fmt.Printf("Category not assigned, ticker %s not found\n", args[0])
	},
}

func init() {
	RootCommand.AddCommand(assignCmd)
}
