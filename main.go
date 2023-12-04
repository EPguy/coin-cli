package main

import (
	"fmt"
	"os"

	cmd "github.com/EPguy/coin-cli/cmd"
	gh "github.com/EPguy/coin-cli/github"
)

func main() {
	cmd.AddCommands()
	gh.InitSingleton()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
