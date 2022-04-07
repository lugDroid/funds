package cmd

import (
	"encoding/json"
	"fmt"
	models "go/funds/Models"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the stocks or funds saved",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Accessing saved stocks/funds data...")
		fmt.Println()

		storedFunds := readFundsFile()

		for _, fund := range storedFunds.Funds {
			fmt.Println("Fund:\t", fund.FundName)
			fmt.Println("Ticker:\t", fund.Ticker)
			fmt.Println("Shares:\t", fund.Shares)
			fmt.Println()
		}
	},
}

func readFundsFile() fundList {
	data, err := ioutil.ReadFile("./funds.json")
	if err != nil {
		fmt.Println(err)
	}

	var fundList fundList

	err = json.Unmarshal(data, &fundList)
	if err != nil {
		fmt.Println(err)
	}

	return fundList
}

type fundList struct {
	Funds []models.Fund
}
