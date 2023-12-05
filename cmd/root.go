package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "coin-cli",
	Short: "CLI that provides coin prices and information",
	Long:  `Made by https://github.com/EPguy`,
}
