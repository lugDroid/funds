package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	"strconv"

	"github.com/spf13/cobra"
)

var category = &cobra.Command{
	Use:   "cat",
	Short: "List the categories defined in the portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing defined categories...")
		fmt.Println()

		categories := data.ReadStorageFile().Categories

		for _, c := range categories {
			fmt.Println("Asset ID:\t", c.Id)
			fmt.Println("Asset Name:\t", c.Name)
			fmt.Println("Allocated:\t", c.PercGoal)
			fmt.Println()
		}
	},
}

var addCategory = &cobra.Command{
	Use:   "add [category name] [allocation value (%)]",
	Short: "Add a new category",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding new category...")
		fmt.Println()

		// Confirm second argument is convertible to float
		percent, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("Number of shares required (float format)")
			return
		}

		// open file, append to list and save
		portfolio := data.ReadStorageFile()

		newCategory := models.Category{
			Id:       len(portfolio.Categories) + 1,
			Name:     args[0],
			PercGoal: float32(percent),
		}

		portfolio.Categories = append(portfolio.Categories, newCategory)
		data.WriteStorageFile(portfolio)
		fmt.Println("New category successfully added")
	},
}

var deleteCategory = &cobra.Command{
	Use:   "delete [category ID]",
	Short: "Delete an existing category",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting existing category...")
		fmt.Println()

		// Confirm argument is convertible to int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Category ID is required")
			return
		}

		portfolio := data.ReadStorageFile()

		for i, c := range portfolio.Categories {
			if c.Id == id {
				portfolio.Categories = append(portfolio.Categories[:i], portfolio.Categories[i+1:]...)
				break
			}
		}

		data.WriteStorageFile(portfolio)

		fmt.Println("Category deleted successfully")
	},
}

func init() {
	RootCommand.AddCommand(category)
	category.AddCommand(addCategory, deleteCategory)
}
