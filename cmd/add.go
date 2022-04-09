package cmd

import (
	"fmt"
	data "go/funds/Data"
	models "go/funds/Models"
	yahoofinanceapi "go/funds/YahooFinanceAPI"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [ticker of stock or fund to add] [number of shares to add]",
	Short: "Add the number of shares of the provided stock or fund",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding new fund %s\n", args[0])
		fmt.Println()

		// Confirm second argument is convertible to float
		shares, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			fmt.Println("Number of shares required (float format)")
			return
		}

		// Get new fund information from yahoo api
		fundData := yahoofinanceapi.GetFundData(args[0])

		newFund := models.Fund{
			FundName: fundData.QuoteResponse.Result[0].LongName,
			Ticker:   args[0],
			Shares:   float32(shares),
		}

		// check if fund is already saved
		fundList := data.ReadFundsFile()

		for _, fund := range fundList.Funds {
			if fund.Ticker == newFund.Ticker {
				fmt.Println("Fund already saved, nothing was added")
				return
			}
		}

		// open file, append and save
		fundList.Funds = append(fundList.Funds, newFund)
		data.WriteFundsFile(fundList)
		fmt.Printf("New stock/fund '%s' successfully saved\n", newFund.FundName)
	},
}

func init() {
	RootCommand.AddCommand(addCmd)
}
