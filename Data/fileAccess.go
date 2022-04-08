package data

import (
	"encoding/json"
	"fmt"
	models "go/funds/Models"
	"io/ioutil"
)

type FundList struct {
	Funds []models.Fund
}

func ReadFundsFile() FundList {
	data, err := ioutil.ReadFile("./funds.json")
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
