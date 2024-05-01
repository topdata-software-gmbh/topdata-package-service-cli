package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var repositoryDetailCmd = &cobra.Command{
	Use:   "detail [repositoryName]",
	Short: "Get the details of a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(repositoryDetailCmd)
}
