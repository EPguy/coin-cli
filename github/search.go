package github

import (
	"log"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

func SearchCoinList() []coinpaprika.Coin {
	coins := GetCoinList()
	var coin []coinpaprika.Coin
	for _, v := range coins {
		if *v.Rank > 0 {
			coin = append(coin, *v)
		}
	}

	return coin
}

func SearchTicker(symbol string) *coinpaprika.Ticker {
	paprikaClient := GetPaprikaClient()
	coins := GetCoinList()
	var coin coinpaprika.Coin
	for _, v := range coins {
		if *v.Symbol == symbol && *v.Rank > 0 {
			coin = *v
			break
		}
	}
	if coin.ID == nil {
		log.Fatalf("Coin is not found with the provided symbol.")
	}
	ticker, err := paprikaClient.Tickers.GetByID(*coin.ID, nil)
	if err != nil {
		log.Fatalf("Error while searching ticker for: %v: %v", symbol, err)
	}
	return ticker
}
