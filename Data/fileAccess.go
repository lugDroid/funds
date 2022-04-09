// Package data implement functions and data types
// used to access the program saved data
package data

import (
	"encoding/json"
	"fmt"
	models "go/funds/Models"
	"io/ioutil"
)

const fundsFilePath string = "./funds.json"

type FundList struct {
	Funds []models.Fund
}

// ReadFundsFile reads the contents of funds.json
// and returns a FundList type
func ReadFundsFile() FundList {
	data, err := ioutil.ReadFile(fundsFilePath)
	if err != nil {
		fmt.Println(err)
	}

	var fundList FundList

	err = json.Unmarshal(data, &fundList)
	if err != nil {
		fmt.Println(err)
	}

	return fundList
}

// WriteFundsFile writes the provided FundList to file
// overwritting any previous content
func WriteFundsFile(list FundList) {
	jsonData, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fundsFilePath, jsonData, 0644)
}
