package yahoofinanceapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAssetData(tick string) YFApiResponse {
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

	var r YFApiResponse

	json.Unmarshal([]byte(body), &r)

	return r
}
