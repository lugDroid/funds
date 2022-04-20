package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var resultCmd = &cobra.Command{
	Use:   "result",
	Short: "Calculate the current value of your portfolio assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calculating portfolio value...")
		fmt.Println()

		results := models.CalculateResults(data.ReadFundsFile())

		writer := tabwriter.NewWriter(os.Stdout, 8, 8, 1, '\t', 0)
		defer writer.Flush()

		fmt.Fprintf(writer, "%s\t%s\t%s\t%s\n", "Asset", "Value", "Category", "% of Total")
		fmt.Fprintf(writer, "--------------------------------------------------\t----------\t----------\t----------\n")

		for _, r := range results {
			fmt.Fprintf(writer, "%s\t%0.2f€\t%s\t%0.2f%%\n", r.Name, r.Value, r.Category, r.PercentOfTotal)
		}
	},
}

func init() {
	RootCommand.AddCommand(resultCmd)
}
