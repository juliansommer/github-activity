package cmd

import (
	"fmt"
	"os"

	"github.com/juliansommer/github-activity/activity"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "GitHub Activity is a CLI tool to analyze user activity on GitHub.",
	Long:  "Github Activity is a CLI tool that allows you to analyze user activity on GitHub. It allows you to fetch user activity by providing a GitHub username.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		events, err := activity.FetchActivity(username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		activity.DisplayActivity(events, username)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
