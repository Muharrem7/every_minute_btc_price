package common

type Asset struct {
	ID       string `json:"id"`
	PriceUsd string `json:"priceUsd"`
	Symbol   string `json:"symbol"`
}

type Response struct {
	Data Asset `json:"data"`
}

type Result struct {
	CoinSymbol string
	CoinName   string
	CoinPrice  string
}
