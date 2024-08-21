package main

import (
	"btc/client"
	"btc/service"
	"github.com/robfig/cron/v3"
	"net/http"
	"time"
)

func main() {

	c := cron.New(cron.WithSeconds())

	cl := &http.Client{
		Timeout: time.Second * 10,
	}

	priceListClient := client.NewPriceListClient(cl, "*****************************")
	srv := service.NewBitcoinService(priceListClient)

	c.AddFunc("0 * * * * *", func() {
		srv.GetPrice()
	})

	c.Start()

	select {}
}
