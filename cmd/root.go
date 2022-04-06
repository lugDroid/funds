package cmd

import "github.com/spf13/cobra"

var RootCommand = &cobra.Command{
	Use:   "funds",
	Short: "Funds is a CLI tool to check stocks and fund values",
}

func init() {
	RootCommand.AddCommand(checkCmd)
	RootCommand.AddCommand(listCmd)
}
