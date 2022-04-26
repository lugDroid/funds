package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	yahoofinanceapi "go/funds/YahooFinanceAPI"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ticker of stock or fund to add] [number of shares to add]",
	Short: "Add the number of shares of the provided stock or fund",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding new asset %s\n", args[0])
		fmt.Println()

		// Confirm second argument is convertible to float
		shares, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("Number of shares required (float format)")
			return
		}

		// Get new asset information from yahoo api
		assetData, err := yahoofinanceapi.GetAssetData(args[0])
		if err != nil {
			fmt.Println("Ticker code not valid")
			os.Exit(1)
		}

		newAsset := models.Asset{
			Name:   assetData.LongName,
			Ticker: args[0],
			Shares: float32(shares),
		}

		// check if asset is already saved
		portfolio := data.ReadStorageFile()

		for _, asset := range portfolio.Assets {
			if asset.Ticker == newAsset.Ticker {
				fmt.Println("Asset already saved, nothing was added")
				return
			}
		}

		// open file, append and save
		portfolio.Assets = append(portfolio.Assets, newAsset)
		data.WriteStorageFile(portfolio)
		fmt.Printf("New asset '%s' successfully saved\n", newAsset.Name)
	},
}

func init() {
	RootCommand.AddCommand(addCmd)
}
