package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var repositoriesListCmd = &cobra.Command{
    Use:   "repositories:list",
    Short: "List all repositories",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing all repositories...")
    },
}

func init() {
    rootCmd.AddCommand(repositoriesListCmd)
}
