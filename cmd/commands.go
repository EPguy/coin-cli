package cmd

import (
	"fmt"
	"math"
	"strings"

	gh "github.com/EPguy/coin-cli/github"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Price = &cobra.Command{
	Use:   "price",
	Short: "Search for the coin price",
	Long:  `coin-cli price <COIN_SYMBOL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := message.NewPrinter(language.English)
		searchVal := strings.Join(args, " ")
		ticker := gh.SearchTicker(searchVal)
		p.Printf("%s price(USD) is : %f", searchVal, *ticker.Quotes["USD"].Price)
	},
}

var List = &cobra.Command{
	Use:   "list",
	Short: "List coins Top 100",
	Long:  `coin-cli list`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		p := message.NewPrinter(language.English)
		tickers := gh.SearchTickerList()[:100]
		for _, ticker := range tickers {
			p.Printf("coin rank : %d\n", *ticker.Rank)
			p.Printf("coin id : %s\n", *ticker.ID)
			p.Printf("coin name : %s\n", *ticker.Name)
			p.Printf("coin symbol : %s\n", *ticker.Symbol)
			p.Printf("coin price : %f$\n", *ticker.Quotes["USD"].Price)
			p.Printf("coin max supply : %d\n", *ticker.MaxSupply)
			p.Printf("coin total supply : %d\n", *ticker.TotalSupply)
			p.Printf("coin market cap : %.0f$\n", math.Floor(*ticker.Quotes["USD"].MarketCap))
			p.Printf("percent change(24h) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange24h)
			p.Printf("percent change(7d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange7d)
			p.Printf("percent change(30d) : %.2f%%\n", *ticker.Quotes["USD"].PercentChange30d)

			fmt.Printf("-----------------------------------------------------\n")
		}
	},
}

func AddCommands() {
	RootCmd.AddCommand(Price, List)
}
