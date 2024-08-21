package client

import (
	"btc/common"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PriceListClient struct {
	Client http.Client
	url    string
}

func NewPriceListClient(c *http.Client, url string) *PriceListClient {
	return &PriceListClient{
		Client: *c,
		url:    url,
	}
}

func (c PriceListClient) GetBitcoinPrice() (*common.Result, error) {

	response, err := http.Get(c.url)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)

	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed with error %s\n", err)
	}

	var apiResponse common.Response
	err = json.Unmarshal(responseData, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("The HTTP request failed to encode  %s\n", err)
	}

	var result common.Result

	coin := apiResponse.Data.ID
	price := apiResponse.Data.PriceUsd
	symbol := apiResponse.Data.Symbol

	result.CoinSymbol = symbol
	result.CoinName = coin
	result.CoinPrice = price

	return &result, nil
}
