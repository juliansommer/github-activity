package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-user-activity",
	Short: "GitHub User Activity is a CLI tool to analyze user activity on GitHub.",
	Long:  "Github User Activity is a command-line tool that allows you to analyze user activity on GitHub. It allows you to fetch user activity by providing a GitHub username.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Github User Activity!")
		if len(args) < 1 {
			fmt.Println("Please provide a GitHub username to analyze.")
			return
		}
		fmt.Println(args[0], "is the GitHub username you want to analyze.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
