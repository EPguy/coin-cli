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

func SearchTickerList() []*coinpaprika.Ticker {
	tickers := GetTickerList()

	return tickers
}

func SearchTicker(symbol string) *coinpaprika.Ticker {
	tickers := GetTickerList()
	var ticker *coinpaprika.Ticker
	for _, v := range tickers {
		if *v.Symbol == symbol && *v.Rank > 0 {
			ticker = v
			break
		}
	}
	if ticker == nil {
		log.Fatalf("Coin is not found with the provided symbol.")
	}
	return ticker
}
