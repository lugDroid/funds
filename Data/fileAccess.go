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

// ReadFundsFile reads the contents of funds.json
// and returns a FundList type
func ReadFundsFile() models.Portfolio {
	data, err := ioutil.ReadFile(fundsFilePath)
	if err != nil {
		fmt.Println(err)
	}

	var fundList models.Portfolio

	err = json.Unmarshal(data, &fundList)
	if err != nil {
		fmt.Println(err)
	}

	return fundList
}

// WriteFundsFile writes the provided FundList to file
// overwritting any previous content
func WriteFundsFile(list models.Portfolio) {
	jsonData, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fundsFilePath, jsonData, 0644)
}
