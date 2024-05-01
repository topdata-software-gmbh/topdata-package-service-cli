package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Commands related to repositories",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all repositories...")
	},
}

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Get details of a repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting details of a repository...")
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
	repositoryCmd.AddCommand(detailsCmd)
	rootCmd.AddCommand(repositoryCmd)
}
