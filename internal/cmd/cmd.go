package main

import (
	"fmt"
	"os"

	"github.com/nc-wittj/plaid-go-legacy/internal/release"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "internal",
		Short: "Internal CLI for development of the plaid-go library",
	}

	rootCmd.AddCommand(releaseCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var releaseCmd = &cobra.Command{
	Use:   "release (major|minor|patch)",
	Short: "prepares the library for releasing a new version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		err := release.Main(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "release.Main: %s\n", err.Error())
			os.Exit(1)
		}
	},
}
