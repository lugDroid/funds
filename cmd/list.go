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

		portfolio := data.ReadStorageFile()

		for _, asset := range portfolio.Assets {
			var categoryName string
			for _, c := range portfolio.Categories {
				if asset.Id == c.Id {
					categoryName = c.Name
					break
				}

				categoryName = "Not Assigned"
			}

			fmt.Println("Fund:\t\t", asset.Name)
			fmt.Println("Ticker:\t\t", asset.Ticker)
			fmt.Println("Category:\t", categoryName)
			fmt.Println("Shares:\t\t", asset.Shares)
			fmt.Println()
		}
	},
}

func init() {
	RootCommand.AddCommand(listCmd)
}
