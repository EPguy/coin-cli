package cmd

import (
	"fmt"
	"strings"

	gh "github.com/EPguy/coin-cli/github"
	"github.com/spf13/cobra"
)

var Price = &cobra.Command{
	Use:   "price",
	Short: "Search for the coin price",
	Long:  `coin-cli price <COIN_SYMBOL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchVal := strings.Join(args, " ")
		ticker := gh.SearchTicker(searchVal)
		fmt.Printf("%s price(USD) is : %f", searchVal, *ticker.Quotes["USD"].Price)
	},
}

var List = &cobra.Command{
	Use:   "list",
	Short: "List coins",
	Long:  `coin-cli list`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tickers := gh.SearchTickerList()
		for _, ticker := range tickers {
			fmt.Printf("coin id : %s\n", *ticker.ID)
			fmt.Printf("coin name : %s\n", *ticker.Name)
			fmt.Printf("coin symbol : %s\n", *ticker.Symbol)
			fmt.Printf("coin price(USD) : %f\n", *ticker.Quotes["USD"].Price)
			fmt.Printf("coin max supply : %d\n", *ticker.MaxSupply)
			fmt.Printf("coin total supply : %d\n", *ticker.TotalSupply)

			fmt.Printf("-----------------------------------------------------\n")
		}
	},
}

func AddCommands() {
	RootCmd.AddCommand(Price, List)
}
