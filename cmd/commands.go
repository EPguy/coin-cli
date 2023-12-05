package cmd

import (
	"fmt"
	"log"
	"math"
	"strconv"
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
	Short: "List coins Top 500",
	Long:  `coin-cli list <MAX_RANK_TO_SHOW>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchVal := strings.Join(args, " ")
		limit, err := strconv.Atoi(searchVal)
		if err != nil {
			log.Fatalf("Please input a number instead of text.")
		}
		tickers := gh.SearchTickerList()
		if limit > len(tickers) {
			limit = len(tickers)
		}
		for _, ticker := range tickers[:limit] {
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
	p.Printf("all time high(ATH) : %f$\n", *ticker.Quotes["USD"].ATHPrice)
	p.Printf("max supply : %d\n", *ticker.MaxSupply)
	p.Printf("total supply : %d\n", *ticker.TotalSupply)
	p.Printf("market cap : %.0f$\n", math.Floor(*ticker.Quotes["USD"].MarketCap))
	p.Printf("percent change(24h) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange24h)
	p.Printf("percent change(7d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange7d)
	p.Printf("percent change(30d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange30d)
}
