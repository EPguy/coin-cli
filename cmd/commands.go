package cmd

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	gh "github.com/EPguy/coin-cli/github"
	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var SortType string

var Info = &cobra.Command{
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
	Short: "Show List coins",
	Long:  `coin-cli list <MAX_RANK_TO_SHOW>`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		searchVal := strings.Join(args, " ")
		maxRankToShow := 2000
		if searchVal != "" {
			var err error
			maxRankToShow, err = strconv.Atoi(searchVal)
			if err != nil {
				log.Fatalf("Please input a number instead of text.")
			}
		}

		tickers := gh.SearchTickerList()
		sortTickerList(tickers)

		if maxRankToShow > len(tickers) {
			maxRankToShow = len(tickers)
		}

		for _, ticker := range tickers[:maxRankToShow] {
			displayTickerInfo(ticker)
			fmt.Printf("-----------------------------------------------------\n")
		}

		fmt.Printf("&&&&&&&&&&&&&&&&&&&&& LIST END &&&&&&&&&&&&&&&&&&&&&\n\n")
	},
}

func AddCommands() {
	RootCmd.PersistentFlags().StringVarP(&SortType, "sort", "s", "rank_asc", "Sort type for list\n(rank_asc, rank_desc, price_asc, price_desc, 1h_asc, 1h_desc, 24h_asc, 24h_desc, 7d_asc, 7d_desc, 30d_asc, 30d_desc)")
	RootCmd.AddCommand(Info, List)
}

func sortTickerList(tickers []*coinpaprika.Ticker) {
	switch SortType {
	case "rank_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Rank > *tickers[j].Rank
		})
	case "price_asc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].Price < *tickers[j].Quotes["USD"].Price
		})
	case "price_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].Price > *tickers[j].Quotes["USD"].Price
		})
	case "1h_asc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange1h < *tickers[j].Quotes["USD"].PercentChange1h
		})
	case "1h_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange1h > *tickers[j].Quotes["USD"].PercentChange1h
		})
	case "24h_asc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange24h < *tickers[j].Quotes["USD"].PercentChange24h
		})
	case "24h_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange24h > *tickers[j].Quotes["USD"].PercentChange24h
		})
	case "7d_asc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange7d < *tickers[j].Quotes["USD"].PercentChange7d
		})
	case "7d_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange7d > *tickers[j].Quotes["USD"].PercentChange7d
		})
	case "30d_asc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange30d < *tickers[j].Quotes["USD"].PercentChange30d
		})
	case "30d_desc":
		sort.Slice(tickers, func(i, j int) bool {
			return *tickers[i].Quotes["USD"].PercentChange30d > *tickers[j].Quotes["USD"].PercentChange30d
		})
	}

}

func displayTickerInfo(ticker *coinpaprika.Ticker) {
	p := message.NewPrinter(language.English)
	p.Printf("rank : %d\n", *ticker.Rank)
	p.Printf("id : %s\n", *ticker.ID)
	p.Printf("name : %s\n", *ticker.Name)
	p.Printf("symbol : %s\n", *ticker.Symbol)
	p.Printf("price : %f$\n", *ticker.Quotes["USD"].Price)
	p.Printf("all time high price(ATH) : %f$\n", *ticker.Quotes["USD"].ATHPrice)
	p.Printf("all time high date(ATH) : %s$\n", *ticker.Quotes["USD"].ATHDate)
	p.Printf("max supply : %d\n", *ticker.MaxSupply)
	p.Printf("total supply : %d\n", *ticker.TotalSupply)
	p.Printf("market cap : %.0f$\n", math.Floor(*ticker.Quotes["USD"].MarketCap))
	p.Printf("percent change(1h) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange1h)
	p.Printf("percent change(24h) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange24h)
	p.Printf("percent change(7d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange7d)
	p.Printf("percent change(30d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange30d)
}
