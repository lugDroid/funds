package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	"strconv"

	"github.com/spf13/cobra"
)

var assets = &cobra.Command{
	Use:   "assets",
	Short: "List the assets defined in the portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing defined assets...")
		fmt.Println()

		assets := data.ReadFundsFile().AssetsAllocation

		for _, asset := range assets {
			fmt.Println("Asset ID:\t", asset.Id)
			fmt.Println("Asset Name:\t", asset.Name)
			fmt.Println("Allocated:\t", asset.Allocated)
			fmt.Println()
		}
	},
}

var addAsset = &cobra.Command{
	Use:   "add [asset name] [allocation value (%)]",
	Short: "Add a new asset",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding new asset...")
		fmt.Println()

		// Confirm second argument is convertible to float
		percent, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("Number of shares required (float format)")
			return
		}

		// open file, append to list and save
		fundList := data.ReadFundsFile()

		newAsset := models.Asset{
			Id:        len(fundList.AssetsAllocation) + 1,
			Name:      args[0],
			Allocated: float32(percent),
		}

		fundList.AssetsAllocation = append(fundList.AssetsAllocation, newAsset)
		data.WriteFundsFile(fundList)
		fmt.Println("New asset successfully added")
	},
}

var deleteAsset = &cobra.Command{
	Use:   "delete [asset ID]",
	Short: "Delete an existing asset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting existing asset...")
		fmt.Println()

		// Confirm argument is convertible to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Asset ID is required")
			return
		}

		fundsList := data.ReadFundsFile()

		for i, asset := range fundsList.AssetsAllocation {
			if asset.Id == id {
				fundsList.AssetsAllocation = append(fundsList.AssetsAllocation[:i], fundsList.AssetsAllocation[i+1:]...)
				break
			}
		}

		data.WriteFundsFile(fundsList)

		fmt.Println("Asset deleted successfully")
	},
}

func init() {
	RootCommand.AddCommand(assets)
	assets.AddCommand(addAsset, deleteAsset)
}
