// Package data implement functions and data types
// used to access the program saved data
package data

import (
	"encoding/json"
	"fmt"
	models "go/funds/Models"
	"io/ioutil"
)

const fundsFilePath string = "./portfolio.json"

// ReadStorageFile reads the contents of funds.json
// and returns a Portfolio type
func ReadStorageFile() models.Portfolio {
	data, err := ioutil.ReadFile(fundsFilePath)
	if err != nil {
		fmt.Println(err)
	}

	var portfolio models.Portfolio

	err = json.Unmarshal(data, &portfolio)
	if err != nil {
		fmt.Println(err)
	}

	return portfolio
}

// WriteStorageFile writes the provided Portfolio to file
// overwritting any previous content
func WriteStorageFile(list models.Portfolio) {
	jsonData, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile(fundsFilePath, jsonData, 0644)
}
