package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all repositories...")
	},
}

func init() {
	repositoryCmd.AddCommand(listCmd)
}
