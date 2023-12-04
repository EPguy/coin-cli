package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Price = &cobra.Command{
	Use:   "price",
	Short: "Search for the coin price",
	Long:  `coin price <COIN_SYMBOL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("hi")
	},
}

var List = &cobra.Command{
	Use:   "list",
	Short: "List coins",
	Long:  `coin list`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("hi")
	},
}

func AddCommands() {
	RootCmd.AddCommand(Price, List)
}
