package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the stocks or funds saved",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Accessing saved stocks/funds data...")
	},
}
