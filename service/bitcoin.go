package service

import (
	"btc/client"
	"fmt"
	"log"
	"time"
)

type BitcoinService struct {
	priceClient *client.PriceListClient
}

func NewBitcoinService(priceClient *client.PriceListClient) *BitcoinService {
	return &BitcoinService{
		priceClient: priceClient,
	}
}

func (b BitcoinService) GetPrice() {
	coin, err := b.priceClient.GetBitcoinPrice()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Coin Symbol: " + coin.CoinSymbol)
	fmt.Println("Coin Name: " + coin.CoinName)
	fmt.Println("Coin Price: " + coin.CoinPrice)
	fmt.Println("Time: " + time.Now().Format("2006-01-02 15:04:05"))
}
