package cmd

import (
	"fmt"
	"math"
	"strings"

	gh "github.com/EPguy/coin-cli/github"
	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Price = &cobra.Command{
	Use:   "info",
	Short: "Search for the coin info",
	Long:  `coin-cli info <COIN_SYMBOL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchVal := strings.Join(args, " ")
		ticker := gh.SearchTicker(searchVal)
		displayTickerInfo(ticker)

	},
}

var List = &cobra.Command{
	Use:   "list",
	Short: "List coins Top 100",
	Long:  `coin-cli list`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tickers := gh.SearchTickerList()[:100]
		for _, ticker := range tickers {
			displayTickerInfo(ticker)
			fmt.Printf("-----------------------------------------------------\n")
		}
	},
}

func AddCommands() {
	RootCmd.AddCommand(Price, List)
}

func displayTickerInfo(ticker *coinpaprika.Ticker) {
	p := message.NewPrinter(language.English)
	p.Printf("rank : %d\n", *ticker.Rank)
	p.Printf("id : %s\n", *ticker.ID)
	p.Printf("name : %s\n", *ticker.Name)
	p.Printf("symbol : %s\n", *ticker.Symbol)
	p.Printf("price : %f$\n", *ticker.Quotes["USD"].Price)
	p.Printf("max supply : %d\n", *ticker.MaxSupply)
	p.Printf("total supply : %d\n", *ticker.TotalSupply)
	p.Printf("market cap : %.0f$\n", math.Floor(*ticker.Quotes["USD"].MarketCap))
	p.Printf("percent change(24h) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange24h)
	p.Printf("percent change(7d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange7d)
	p.Printf("percent change(30d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange30d)
}
