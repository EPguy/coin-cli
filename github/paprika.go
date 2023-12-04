package github

import (
	"sync"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

var once sync.Once
var paprikaClient *coinpaprika.Client
var coinList []*coinpaprika.Coin

func InitSingleton() {
	once.Do(func() {
		paprikaClient = coinpaprika.NewClient(nil)
		coinList, _ = paprikaClient.Coins.List()
	})
}

func GetPaprikaClient() *coinpaprika.Client {
	return paprikaClient
}

func GetCoinList() []*coinpaprika.Coin {
	return coinList
}
