package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var repositoryDetailsCmd = &cobra.Command{
	Use:   "details [repositoryName]",
	Short: "Get the details of a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s\n", args[0])
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryDetailsCmd)
}
