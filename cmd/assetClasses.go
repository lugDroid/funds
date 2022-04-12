package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	"strconv"

	"github.com/spf13/cobra"
)

var assetClasses = &cobra.Command{
	Use:   "asset-classes",
	Short: "List the asset classes defined in the portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing defined asset classes...")
		fmt.Println()

		assetClasses := data.ReadFundsFile().AssetClasses

		for _, class := range assetClasses {
			fmt.Println("Class ID:\t\t", class.Id)
			fmt.Println("Class Name:\t\t", class.Name)
			fmt.Println("Percent of Total:\t", class.PercentOfTotal)
			fmt.Println()
		}
	},
}

var addClass = &cobra.Command{
	Use:   "add [class name] [percent of total portfolio value]",
	Short: "Add a new asset class",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding new asset class...")
		fmt.Println()

		// Confirm second argument is convertible to float
		percent, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("Number of shares required (float format)")
			return
		}

		// open file, append to list and save
		fundList := data.ReadFundsFile()

		newClass := models.AssetClass{
			Id:             len(fundList.AssetClasses) + 1,
			Name:           args[0],
			PercentOfTotal: float32(percent),
		}

		fundList.AssetClasses = append(fundList.AssetClasses, newClass)
		data.WriteFundsFile(fundList)
		fmt.Println("New asset class successfully added")
	},
}

var deleteClass = &cobra.Command{
	Use:   "delete [class ID]",
	Short: "Delete an existing asset class",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting existing asset class...")
		fmt.Println()

		// Confirm argument is convertible to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Asset class ID is required")
			return
		}

		fundsList := data.ReadFundsFile()

		for i, assetClass := range fundsList.AssetClasses {
			if assetClass.Id == id {
				fundsList.AssetClasses = append(fundsList.AssetClasses[:i], fundsList.AssetClasses[i+1:]...)
				break
			}
		}

		data.WriteFundsFile(fundsList)

		fmt.Println("Asset class deleted successfully")
	},
}

func init() {
	RootCommand.AddCommand(assetClasses)
	assetClasses.AddCommand(addClass, deleteClass)
}
