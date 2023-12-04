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
	Long:  `coin price <COIN_SYMBOL>`,
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
	Long:  `coin list`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		coins := gh.SearchCoinList()
		for _, coin := range coins {
			ticker := gh.SearchTicker(*coin.Symbol)
			fmt.Printf("coin id : %s\n", *coin.ID)
			fmt.Printf("coin name : %s\n", *coin.Name)
			fmt.Printf("coin symbol : %s\n", *coin.Symbol)
			fmt.Printf("coin price(USD) : %f", *ticker.Quotes["USD"].Price)
		}
	},
}

func AddCommands() {
	RootCmd.AddCommand(Price, List)
}
