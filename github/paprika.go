package github

import (
	"sync"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

var once sync.Once
var paprikaClient *coinpaprika.Client
var coinList []*coinpaprika.Coin
var tickerList []*coinpaprika.Ticker

func InitSingleton() {
	once.Do(func() {
		paprikaClient = coinpaprika.NewClient(nil)
		coinList, _ = paprikaClient.Coins.List()
		tickerList, _ = paprikaClient.Tickers.List(nil)
	})
}

func GetPaprikaClient() *coinpaprika.Client {
	return paprikaClient
}

func GetCoinList() []*coinpaprika.Coin {
	return coinList
}

func GetTickerList() []*coinpaprika.Ticker {
	return tickerList
}
