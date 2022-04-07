package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ticker of stock or fund to add] [number of shares to add]",
	Short: "Add the number of shares of the provided stock or fund",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding new fund %s\n", args[0])
	},
}
