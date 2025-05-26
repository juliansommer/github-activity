package cmd

import (
	"fmt"
	"os"

	"github.com/juliansommer/github-activity/activity"
	"github.com/spf13/cobra"
)

var perPage int
var pageNum int

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "GitHub Activity is a CLI tool to analyze user activity on GitHub.",
	Long:  "Github Activity is a CLI tool that allows you to analyze user activity on GitHub. It allows you to fetch user activity by providing a GitHub username.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		events, err := activity.FetchActivity(username, perPage, pageNum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		activity.DisplayActivity(events, username)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&perPage, "per_page", "n", 10, "Number of events to fetch per page")
	rootCmd.Flags().IntVarP(&pageNum, "page", "p", 1, "Page number to fetch")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
