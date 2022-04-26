package yahoofinanceapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAssetData(tick string) (*Result, error) {
	baseUrl := "https://yfapi.net/v6/finance/quote?region=ES&lang=en&symbols="
	baseUrl += tick
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		println(err)
	}

	apiKey := os.Getenv("YF_API_KEY")
	if apiKey == "" {
		fmt.Println("API key environment variable not found")
		os.Exit(1)
	}

	req.Header.Set("x-api-key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}

	var response YFApiResponse

	json.Unmarshal([]byte(body), &response)

	if len(response.QuoteResponse.Result) == 0 {
		return nil, errors.New("no results received")
	}

	return &response.QuoteResponse.Result[0], nil
}
