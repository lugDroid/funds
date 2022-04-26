package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ticker of stock or fund to delete]",
	Short: "Delete the provided stock or fund",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Deleting asset '%s'\n", args[0])
		fmt.Println()

		assets := data.ReadStorageFile().Assets

		for i, a := range assets {
			if a.Ticker == args[0] {
				assets = append(assets[:i], assets[i+1:]...)
				break
			}
		}

		data.WriteStorageFile(models.Portfolio{
			Assets: assets,
		})

		fmt.Println("Fund/stock deleted successfully")
	},
}

func init() {
	RootCommand.AddCommand(deleteCmd)
}
