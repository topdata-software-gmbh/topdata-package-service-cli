package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var repositoryDetailCmd = &cobra.Command{
    Use:   "repository detail",
    Short: "Show details of a repository",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Showing details of a repository...")
    },
}

func init() {
    rootCmd.AddCommand(repositoryDetailCmd)
}
