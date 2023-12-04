package main

import (
	"fmt"
	"os"

	"github.com/EPguy/coin-cli/cmd"
)

func main() {
	cmd.AddCommands()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
