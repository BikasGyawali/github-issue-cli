package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-issue-cli",
	Short: "A CLI tool for managing GitHub issues",
	Long:  `A command-line application to create, list, and close GitHub issues using the GitHub API.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
